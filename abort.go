package piscine

func Abort(a, b, c, d, e int) int {
	sort := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if sort[i] > sort[j] {
				sort[i], sort[j] = sort[j], sort[i]
			}
		}
	}
	return sort[2]
}
