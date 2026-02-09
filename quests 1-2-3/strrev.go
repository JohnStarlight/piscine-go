package piscine

func StrRev(s string) string {
	runes := []rune(s)
	n := 0
	for range runes {
		n++
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
