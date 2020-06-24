package mymath

func MySum(a int, b ...int) int {
	sum := a
	if len(b) != 0 {
		for _, k := range b {
			sum += k
		}
	}
	return sum
}
