package generate

import (
	"sync"

	"github.com/EduardGomezEscandell/GoMandelbrot/colors"
	"github.com/EduardGomezEscandell/GoMandelbrot/frames"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

type Config struct {
	Width           uint
	Height          uint
	Cmap            colors.Colormap
	Maxiter         uint
	Xspan           [2]float64
	Yspan           [2]float64
	OutputFilename  string
	MetaData        string
	Verbosity       bool
	SubscalingLevel uint
}

type IterationCounter = func(complex128, uint) uint

func (self *Config) DefineComplexFrame(center complex128, real_span float64, imag_span float64) {
	self.Xspan = [2]float64{real(center) - real_span/2.0, real(center) + real_span/2.0}
	self.Yspan = [2]float64{imag(center) - imag_span/2.0, imag(center) + imag_span/2.0}
}

// Standard pixels
func generate_pixel(row uint, col uint, config *Config, f IterationCounter) int {
	c := maths.PixelToCoordinate(int(row), int(col), config.Width, config.Height, config.Xspan, config.Yspan)
	return int(f(c, config.Maxiter))
}

func generate_row(row frames.Row[int], config *Config, f IterationCounter) {
	frames.ForEachIndexed(row.Begin(), row.End(), func(iter *int, col uint) {
		*iter = generate_pixel(row.Index(), col, config, f)
	})
}

// Subsampled pixels
func generate_pixel_subsampled(row uint, col uint, config *Config, f IterationCounter) int {
	bottom_left := maths.PixelToCoordinate(int(row)-1, int(col)-1, config.Width, config.Height, config.Xspan, config.Yspan)
	top_right := maths.PixelToCoordinate(int(row)+1, int(col)+1, config.Width, config.Height, config.Xspan, config.Yspan)
	center := maths.PixelToCoordinate(int(row), int(col), config.Width, config.Height, config.Xspan, config.Yspan)

	bl_boundary := (bottom_left + center) / 2.0
	tr_boundary := (top_right + center) / 2.0

	delta := (tr_boundary - bl_boundary) / complex(float64(config.SubscalingLevel), 0)

	count := uint(0)
	for i := uint(0); i < config.SubscalingLevel; i++ {
		real_ := real(bl_boundary) + float64(i)*real(delta)
		for j := uint(0); j < config.SubscalingLevel; j++ {
			imag_ := imag(bl_boundary) + float64(j)*imag(delta)
			c := complex(real_, imag_)
			count += f(c, config.Maxiter)
		}
	}
	return int(float64(count) / float64(config.SubscalingLevel*config.SubscalingLevel))
}

func generate_row_subsampled(row frames.Row[int], config *Config, f IterationCounter) {
	frames.ForEachIndexed(row.Begin(), row.End(), func(iter *int, col uint) {
		*iter = generate_pixel_subsampled(row.Index(), col, config, f)
	})
}

// Big loops

// Sequential version, for debugging
func generate_sequential(config *Config, f IterationCounter) frames.IntFrame {
	if config.SubscalingLevel == 1 {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			generate_row(frame.GetRow(i), config, f)
		}
		return frame
	} else {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			generate_row_subsampled(frame.GetRow(i), config, f)
		}
		return frame
	}
}

// Concurrent version, for production
func generate_concurrent(config *Config, f IterationCounter) frames.IntFrame {
	var wg sync.WaitGroup
	defer wg.Wait() // This is a barrier

	if config.SubscalingLevel == 1 {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			wg.Add(1)
			go func(idx uint) {
				defer wg.Done()
				generate_row(frame.GetRow(idx), config, f)
			}(i)
		}
		return frame
	} else {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			wg.Add(1)
			go func(idx uint) {
				defer wg.Done()
				generate_row_subsampled(frame.GetRow(idx), config, f)
			}(i)
		}
		return frame
	}
}

var Generate = generate_concurrent
