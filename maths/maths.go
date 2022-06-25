package maths

import (
	"github.com/EduardGomezEscandell/GoMandelbrot/imageIO"
)

type Complex struct {
	Real float64
	Imag float64
}

func rescale(xi_min int, xi int, xi_max int, x_min float64, x_max float64) float64 {
	return float64(xi-xi_min)/float64(xi_max-xi_min)*(x_max-x_min) + x_min
}

func PixelToCoordinate(image *imageIO.SequentialImage, xspan [2]float64, yspan [2]float64) Complex {
	row, col := image.Position()
	real := rescale(0, col, image.Width, xspan[0], xspan[1])
	imag := rescale(0, image.Height-row, image.Height, yspan[0], yspan[1])
	return Complex{real, imag}
}

func mandelbrot_iteration(z Complex, c Complex) Complex {
	return Complex{
		Real: z.Real*z.Real - z.Imag*z.Imag + c.Real,
		Imag: 2*z.Real*z.Imag + c.Imag,
	}
}

func (z *Complex) MagnitudeSquared() float64 {
	return z.Real*z.Real + z.Imag*z.Imag
}

func MandelbrotDivergenceIter(c Complex, max_iter int) int {
	z := Complex{0, 0}
	for i := 0; i < max_iter; i++ {
		z = mandelbrot_iteration(z, c)
		if z.MagnitudeSquared() > 4 {
			return i
		}
	}
	return max_iter
}
