package colors

var ct_bright Colortable = Colortable{
	lower_bounds: []float64{0, 0.2, 0.4, 0.6, 0.8, 1},
	colors: []Color{
		ColorFromHex(0xff0000),
		ColorFromHex(0x00ff00),
		ColorFromHex(0x0000ff),
		ColorFromHex(0xffff00),
		ColorFromHex(0x00ffff),
		ColorFromHex(0xff00ff)},
	max_color: ColorFromHex(0),
	len:       6,
}

func multicolorEval(cmap *Colormap, value int) Color {

	if value == cmap.Upper {
		return ct_bright.max_color
	}
	d := value % ct_bright.len
	x := float64(d) / float64(ct_bright.len)
	// fmt.Printf("%d -> %d -> %f\n", value, d, x)
	return ct_bright.Get(x)
}
