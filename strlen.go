package piscine

func StrLen(s string) int {
	c := 0
	for i := 0; i < len(s)-1; i++ {
		c++
	}
	return c
}
