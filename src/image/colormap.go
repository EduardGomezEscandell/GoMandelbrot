package image

import (
	"errors"
	"fmt"
	"math"
)

type Colormap struct {
	Lower     int
	Upper     int
	eval      func(*Colormap, float64) Color
	nlp       float64 // Nonlinear parameter
	root_nlp  float64 // Square root of nonlinear parameter
	denom_nlp float64 // Denominator of nonlinear parameter
}

func (self *Colormap) value_nonlinearization(x float64) float64 {
	/*
	 Cached version of:
	   sqrt(x + nlp) - sqrt(nlp) / (sqrt(1.0 + nlp) - sqrt(nlp))
	*/
	return (math.Sqrt(x+self.nlp) - self.root_nlp) / self.denom_nlp
}

func clamp_normalized(value float64) float64 {
	return math.Max(0, math.Min(value, 1))
}

func normalize(xi_min int, xi int, xi_max int) float64 {
	return clamp_normalized(float64(xi-xi_min) / float64(xi_max-xi_min))
}

func (self *Colormap) Eval(value int) Color {
	value_normalized := normalize(self.Lower, value, self.Upper)
	x := self.value_nonlinearization(value_normalized)
	return self.eval(self, x)
}

func grayScaleEval(cmap *Colormap, value_normalized float64) Color {
	x := 255 - uint8(value_normalized*255)
	return Color{
		R: x,
		G: x,
		B: x}
}

func rainbowEval(cmap *Colormap, value_normalized float64) Color {
	x := value_normalized * 360 * 0.5
	return ColorFromHSV(x, 0.8, 1)
}

func ColormapFactory(name string, lower int, upper int, nonlinear_parameter float64) Colormap {
	nlp := nonlinear_parameter
	root_nlp := math.Sqrt(nlp)
	denom_nlp := math.Sqrt(1.0+nlp) - root_nlp

	if name == "grayscale" {
		return Colormap{lower, upper, grayScaleEval, nlp, root_nlp, denom_nlp}
	} else if name == "rainbow" {
		return Colormap{lower, upper, rainbowEval, nlp, root_nlp, denom_nlp}
	}
	panic(errors.New(fmt.Sprintf("Colormap %s not found", name)))
}
