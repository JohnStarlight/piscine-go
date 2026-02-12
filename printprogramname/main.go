package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	// Find last slash
	lastSlash := 0
	for i, r := range name {
		if r == '/' {
			lastSlash = i + 1
		}
	}

	// Print everything after the last slash
	for _, r := range name[lastSlash:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
