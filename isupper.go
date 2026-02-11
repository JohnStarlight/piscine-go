package piscine

func IsUpper(s string) bool {
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if !(r[i] >= 'A' && r[i] <= 'Z') {
			return false
		}
	}
	return true
}
