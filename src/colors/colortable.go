package colors

type Colortable struct {
	lower_bounds []float64 // Must be sorted in ascending order
	colors       []Color
	max_color    Color
	len          int // Must be computed
}

func (self *Colortable) binary_search(value float64, lower int, upper int) int {
	// fmt.Printf("%f <= %f <= %f\n", self.lower_bounds[lower], value, self.lower_bounds[upper])

	if lower == upper {
		return lower
	}

	middle_idx := (lower + upper) / 2
	middle_value := self.lower_bounds[middle_idx]
	if value <= middle_value {
		return self.binary_search(value, lower, middle_idx)
	} else {
		return self.binary_search(value, middle_idx+1, upper)
	}
}

func (self *Colortable) Get(value_normalized float64) Color {
	// Expects a value beteen 0 and 1
	idx := self.binary_search(float64(value_normalized), 0, self.len-1)
	return self.colors[idx]
}
