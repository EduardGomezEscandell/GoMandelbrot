package generate

import (
	"sync"

	"github.com/EduardGomezEscandell/GoMandelbrot/colors"
	"github.com/EduardGomezEscandell/GoMandelbrot/frames"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

type Config struct {
	Width            uint
	Height           uint
	Cmap             colors.Colormap
	Maxiter          int
	Xspan            [2]float64
	Yspan            [2]float64
	OutputFilename   string
	MetaData         string
	Verbosity        bool
	SubscalingFactor uint
}

func (self *Config) DefineComplexFrame(center maths.Complex, real_span float64, imag_span float64) {
	self.Xspan = [2]float64{center.Real - real_span/2, center.Real + real_span/2}
	self.Yspan = [2]float64{center.Imag - imag_span/2, center.Imag + imag_span/2}
}

// Standard pixels
func generate_pixel(row uint, col uint, config *Config) int {
	c := maths.PixelToCoordinate(row, col, config.Width, config.Height, config.Xspan, config.Yspan)
	return maths.MandelbrotDivergenceIter(c, config.Maxiter)
}

func generate_row(row frames.Row[int], config *Config) {
	frames.ForEachIndexed(row.Begin(), row.End(), func(iter *int, col uint) {
		*iter = generate_pixel(row.Index(), col, config)
	})
}

// Subsampled pixels
func generate_pixel_subsampled(row uint, col uint, config *Config) int {
	bottom_left := maths.PixelToCoordinate(row-1, col-1, config.Width, config.Height, config.Xspan, config.Yspan)
	top_right := maths.PixelToCoordinate(row+1, col+1, config.Width, config.Height, config.Xspan, config.Yspan)
	center := maths.PixelToCoordinate(row, col, config.Width, config.Height, config.Xspan, config.Yspan)

	bl_boundary := bottom_left.Add(center).DivideScalar(2.0)
	tr_boundary := top_right.Add(center).DivideScalar(2.0)

	delta := tr_boundary.Subtract(bl_boundary).DivideScalar(float64(config.SubscalingFactor))

	count := 0
	c := maths.Complex{}
	for i := uint(0); i < config.SubscalingFactor; i++ {
		c.Real = bl_boundary.Real + float64(i)*delta.Real
		for j := uint(0); j < config.SubscalingFactor; j++ {
			c.Imag = bl_boundary.Imag + float64(j)*delta.Imag
			count += maths.MandelbrotDivergenceIter(c, config.Maxiter)
		}
	}
	return int(float64(count) / float64(config.SubscalingFactor*config.SubscalingFactor))
}

func generate_row_subsampled(row frames.Row[int], config *Config) {
	frames.ForEachIndexed(row.Begin(), row.End(), func(iter *int, col uint) {
		*iter = generate_pixel_subsampled(row.Index(), col, config)
	})
}

// Big loops

// Sequential version, for debugging
func generate_sequential(config *Config) frames.IntFrame {
	if config.SubscalingFactor == 1 {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			generate_row(frame.GetRow(i), config)
		}
		return frame
	} else {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			generate_row_subsampled(frame.GetRow(i), config)
		}
		return frame
	}
}

// Concurrent version, for production
func generate_concurrent(config *Config) frames.IntFrame {
	var wg sync.WaitGroup
	defer wg.Wait() // This is a barrier

	if config.SubscalingFactor == 1 {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			wg.Add(1)
			go func(idx uint) {
				defer wg.Done()
				generate_row(frame.GetRow(idx), config)
			}(i)
		}
		return frame
	} else {
		frame := frames.NewIntFrame(config.Width, config.Height)
		for i := uint(0); i < config.Height; i++ {
			wg.Add(1)
			go func(idx uint) {
				defer wg.Done()
				generate_row_subsampled(frame.GetRow(idx), config)
			}(i)
		}
		return frame
	}
}

var Generate = generate_concurrent
