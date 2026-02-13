package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	n := max - min
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = min + i
	}
	return s
}
