package piscine

func Unmatch(a []int) int {
	for _, v1 := range a {
		cnt := 0
		for _, v2 := range a {
			if v1 == v2 {
				cnt++
			}
		}
		if cnt%2 != 0 {
			return v1
		}
	}
	return -1
}
