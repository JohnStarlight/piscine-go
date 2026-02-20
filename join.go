package piscine

func Join(strs []string, sep string) string {
	var res string
	for i, v := range strs {
		if i == len(strs)-1 {
			res += v
		} else {
			res += v + sep
		}
	}
	return res
}
