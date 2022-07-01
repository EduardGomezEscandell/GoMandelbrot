package maths

import (
	"fmt"
	"strings"
)

type Complex struct {
	Real float64
	Imag float64
}

func rescale(xi_min uint, xi int, xi_max uint, x_min float64, x_max float64) float64 {
	return float64(xi-int(xi_min))/float64(xi_max-xi_min)*(x_max-x_min) + x_min
}

func PixelToCoordinate(row int, col int, width uint, height uint, xspan [2]float64, yspan [2]float64) Complex {
	real := rescale(0, col, width, xspan[0], xspan[1])
	imag := rescale(0, int(height)-row, height, yspan[0], yspan[1])
	return Complex{real, imag}
}

func mandelbrot_iteration(z Complex, c Complex) Complex {
	return Complex{
		Real: z.Real*z.Real - z.Imag*z.Imag + c.Real,
		Imag: 2*z.Real*z.Imag + c.Imag,
	}
}

func (z Complex) Add(x Complex) Complex {
	x.Real = z.Real + x.Real
	x.Imag = z.Imag + x.Imag
	return x
}

func (z Complex) Subtract(x Complex) Complex {
	x.Real = z.Real - x.Real
	x.Imag = z.Imag - x.Imag
	return x
}

func (z Complex) DivideScalar(n float64) Complex {
	z.Real = z.Real / n
	z.Imag = z.Imag / n
	return z
}

func (z *Complex) MagnitudeSquared() float64 {
	return z.Real*z.Real + z.Imag*z.Imag
}

func MandelbrotDivergenceIter(c Complex, max_iter uint) uint {
	z := Complex{0, 0}
	for i := uint(0); i < max_iter; i++ {
		z = mandelbrot_iteration(z, c)
		if z.MagnitudeSquared() > 4 {
			return i
		}
	}
	return max_iter
}

func ParseComplex(name string) Complex {
	name_trimmed := strings.Replace(name, " ", "", -1)
	var c Complex

	// Canonical form
	if found, err := fmt.Sscanf(name_trimmed, "%f%fi", &c.Real, &c.Imag); err != nil && found == 2 {
		return c
	}

	// Real only
	if found, err := fmt.Sscanf(name_trimmed, "%f", &c.Real, &c.Imag); err != nil && found == 1 {
		return c
	}

	// Imaginary only
	if found, err := fmt.Sscanf(name_trimmed, "%fi", &c.Real, &c.Imag); err != nil && found == 1 {
		return c
	}

	panic(fmt.Sprintf("Failed to parse complex number: %s", name_trimmed))
}

func Swap(a *int, b *int) {
	tmp := *b
	*b = *a
	*a = tmp
}
