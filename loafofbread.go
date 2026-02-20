package piscine

func LoafOfBread(str string) string {
	// count non-space chars
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0

	for i < len(str) {
		collected := 0

		// collect 5 non-space chars
		for i < len(str) && collected < 5 {
			if str[i] != ' ' {
				result += string(str[i])
				collected++
			}
			i++
		}

		if collected < 5 {
			break
		}

		// skip next character (even space)
		if i < len(str) {
			i++
		}
	}

	return result + "\n"
}
