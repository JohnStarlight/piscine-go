package piscine

func IsLower(s string) bool {
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if !(r[i] >= 'a' && r[i] <= 'z') {
			return false
		}
	}
	return true
}
