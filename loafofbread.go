package piscine

func LoafOfBread(str string) string {
	// count non-space chars
	n := 0
	for _, r := range str {
		if r != ' ' {
			n++
		}
	}
	if n < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0

	for i < len(str) {
		collected := 0

		// collect exactly 5 non-space chars
		for i < len(str) && collected < 5 {
			result += string(str[i])
			if str[i] != ' ' {
				collected++
			}
			i++
		}

		if collected < 5 {
			break
		}

		// skip next char completely (don't print it)
		if i < len(str) {
			i++
		}
	}

	return result + "\n"
}
