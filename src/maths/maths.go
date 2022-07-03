package maths

func rescale(xi_min uint, xi int, xi_max uint, x_min float64, x_max float64) float64 {
	return float64(xi-int(xi_min))/float64(xi_max-xi_min)*(x_max-x_min) + x_min
}

func PixelToCoordinate(row int, col int, width uint, height uint, xspan [2]float64, yspan [2]float64) complex128 {
	real := rescale(0, col, width, xspan[0], xspan[1])
	imag := rescale(0, int(height)-row, height, yspan[0], yspan[1])
	return complex(real, imag)
}

func MandelbrotDivergencePeriod(c complex128, max_iter uint) uint {
	z := complex128(0)
	for i := uint(0); i < max_iter; i++ {
		z = z*z + c
		abs2 := real(z)*real(z) + imag(z)*imag(z)
		if abs2 > 4.0 {
			return i
		}
	}
	return max_iter
}

func JuliaDivergencePeriod(z complex128, a complex128, max_iter uint) uint {
	for i := uint(0); i < max_iter; i++ {
		z = z*z + a
		abs2 := real(z)*real(z) + imag(z)*imag(z)
		if abs2 > 4.0 {
			return i
		}
	}
	return max_iter
}
