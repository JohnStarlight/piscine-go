package piscine

func IsPrintable(s string) bool {
	r := []rune(s)

	for i := 0; i < len(r); i++ {
		if r[i] < 32 || r[i] > 126 {
			return false
		}
	}
	return true
}
