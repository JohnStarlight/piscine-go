package piscine

func StrLen(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		c++
	}
	return c
}
