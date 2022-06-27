package generate

import (
	"github.com/EduardGomezEscandell/GoMandelbrot/colors"
	"github.com/EduardGomezEscandell/GoMandelbrot/frames"
	"github.com/EduardGomezEscandell/GoMandelbrot/maths"
)

type Config struct {
	Width          uint
	Height         uint
	Cmap           colors.Colormap
	Maxiter        int
	Xspan          [2]float64
	Yspan          [2]float64
	OutputFilename string
	MetaData       string
	Verbosity      bool
}

func (self *Config) DefineComplexFrame(center maths.Complex, real_span float64, imag_span float64) {
	self.Xspan = [2]float64{center.Real - real_span/2, center.Real + real_span/2}
	self.Yspan = [2]float64{center.Imag - imag_span/2, center.Imag + imag_span/2}
}

func generate_row(row frames.Row[int], config *Config) {
	for it := row.Begin(); it != row.End(); it.Next() {
		row, col := it.Position()
		c := maths.PixelToCoordinate(row, col, config.Width, config.Height, config.Xspan, config.Yspan)
		*it.Ptr() = maths.MandelbrotDivergenceIter(c, config.Maxiter)
	}
}

func GenerateSequential(config *Config) frames.IntFrame {
	frame := frames.NewIntFrame(config.Width, config.Height)
	for i := uint(0); i < config.Height; i++ {
		generate_row(frame.GetRow(i), config)
	}
	return frame
}

func GenerateConcurrent(config *Config) frames.IntFrame {
	frame := frames.NewIntFrame(config.Width, config.Height)
	for i := uint(0); i < config.Height; i++ {
		go generate_row(frame.GetRow(i), config)
	}
	return frame
}
