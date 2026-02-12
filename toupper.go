package piscine

func ToUpper(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] <= 122 && r[i] >= 97 {
			r[i] = r[i] - 32
		}
	}
	return string(r)
}
