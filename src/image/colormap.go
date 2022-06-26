package image

import (
	"errors"
	"fmt"
	"math"
)

func clamp_normalized(value float64) float64 {
	return math.Max(0, math.Min(value, 1))
}

func normalize(xi_min int, xi int, xi_max int) float64 {
	return clamp_normalized(float64(xi-xi_min) / float64(xi_max-xi_min))
}

type Colormap struct {
	Lower int
	Upper int
	eval  func(*Colormap, float64) Color
}

func (cmap *Colormap) Eval(value int) Color {
	value_normalized := normalize(cmap.Lower, value, cmap.Upper)
	return cmap.eval(cmap, value_normalized)
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

func ColormapFactory(name string, lower int, upper int) Colormap {
	if name == "grayscale" {
		return Colormap{lower, upper, grayScaleEval}
	} else if name == "rainbow" {
		return Colormap{lower, upper, rainbowEval}
	}
	panic(errors.New(fmt.Sprintf("Colormap %s not found", name)))
}
