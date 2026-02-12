package piscine

func ToLower(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= 65 && r[i] <= 90 {
			r[i] = r[i] + 32
		}
	}
	return string(r)
}
