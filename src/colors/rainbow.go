package colors

func rainbowEval(cmap *Colormap, value int) Color {
	value_normalized := normalize(cmap.Lower, value, cmap.Upper)
	value_normalized = cmap.value_nonlinearization(value_normalized)
	x := value_normalized * 360 * 0.5
	return ColorFromHSV(x, 0.8, 1)
}
