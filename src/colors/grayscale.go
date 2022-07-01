package colors

func grayScaleEval(cmap *Colormap, value int) Color {
	value_normalized := normalize(cmap.Lower, value, cmap.Upper)
	value_normalized = cmap.value_nonlinearization(value_normalized)

	x := 255 - uint8(value_normalized*255)
	return Color{
		R: x,
		G: x,
		B: x}
}
