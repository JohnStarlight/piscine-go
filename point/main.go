package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	a := n / 10
	b := n % 10
	if a > 0 {
		z01.PrintRune(rune(a + 48))
	}
	z01.PrintRune(rune(b + 48))
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(points.y)
	z01.PrintRune('\n')
}
