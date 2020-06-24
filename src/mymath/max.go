package mymath

func MyMax(a int, b ...int) int {
	max := a
	if len(b) != 0 {
		for _, k := range b {
			if k>max{
				max=k
			}
		}
	}
	return max
}