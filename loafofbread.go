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

		// collect exactly 5 non-space chars (skip spaces)
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

		// skip next non-space char
		for i < len(str) {
			if str[i] != ' ' {
				i++
				break
			}
			i++
		}
	}

	return result + "\n"
}