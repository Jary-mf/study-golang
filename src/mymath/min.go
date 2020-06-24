package mymath

func MyMin(a int, b ...int) int {
	min := a
	if len(b) != 0 {
		for _, k := range b {
			if k < min {
				min = k
			}
		}
	}
	return min
}
