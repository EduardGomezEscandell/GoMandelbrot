package imageIO

import "math"

type Color struct {
	R uint8
	G uint8
	B uint8
}

func ColorFromHex(hex int32) Color {
	return Color{
		R: uint8((hex >> 16) & 0xFF),
		G: uint8((hex >> 8) & 0xFF),
		B: uint8(hex & 0xFF),
	}
}

func ColorFromHSV(h float64, s float64, v float64) Color {
	C := v * s
	hextant := uint8(h / 60.0)
	X := C * (1.0 - math.Abs(float64(hextant%2.0)-1.0))

	m := v - C
	x := uint8((X + m) * 255)
	c := uint8((C + m) * 255)
	if hextant < 1.0 {
		return Color{c, x, 0}
	} else if hextant < 2.0 {
		return Color{x, c, 0}
	} else if hextant < 3.0 {
		return Color{0, c, x}
	} else if hextant < 4.0 {
		return Color{0, x, c}
	} else if hextant < 5.0 {
		return Color{x, 0, c}
	} else {
		return Color{c, 0, x}
	}
}
