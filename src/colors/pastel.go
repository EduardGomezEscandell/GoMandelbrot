package colors

// Colorscheme from https://www.schemecolor.com/therapist.php

var ct_pastel Colortable = Colortable{
	lower_bounds: []float64{0, 0.2, 0.4, 0.6, 0.8, 1},
	colors: []Color{
		ColorFromHex(0xA3CEB1),
		ColorFromHex(0xEBEBD3),
		ColorFromHex(0xE8D3B6),
		ColorFromHex(0xA3AEC0),
		ColorFromHex(0xE0BCC6)},
	max_color: ColorFromHex(0),
	len:       5,
}

func pastelEval(cmap *Colormap, value int) Color {
	if value == cmap.Upper {
		return ct_pastel.max_color
	}
	d := value % ct_pastel.len
	x := float64(d) / float64(ct_pastel.len)

	return ct_pastel.Get(x)
}
