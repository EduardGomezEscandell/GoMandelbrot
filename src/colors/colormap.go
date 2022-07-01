package colors

import (
	"errors"
	"fmt"
	"math"
)

type Colormap struct {
	Lower     int
	Upper     int
	eval      func(*Colormap, int) Color
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
	return self.eval(self, value)
}

func ColormapFactory(name string, lower int, upper int, nonlinear_parameter float64) Colormap {
	nlp := nonlinear_parameter
	root_nlp := math.Sqrt(nlp)
	denom_nlp := math.Sqrt(1.0+nlp) - root_nlp

	if name == "grayscale" {
		return Colormap{lower, upper, grayScaleEval, nlp, root_nlp, denom_nlp}
	} else if name == "rainbow" {
		return Colormap{lower, upper, rainbowEval, nlp, root_nlp, denom_nlp}
	} else if name == "multicolor" {
		return Colormap{lower, upper, multicolorEval, 0, 0, 0}
	} else if name == "pastel" {
		return Colormap{lower, upper, pastelEval, nlp, root_nlp, denom_nlp}
	}
	panic(errors.New(fmt.Sprintf("Colormap %s not found", name)))
}
