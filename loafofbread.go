package piscine

func LoafOfBread(str string) string {
	// Remove spaces
	var chars []rune
	for _, c := range str {
		if c != ' ' {
			chars = append(chars, c)
		}
	}

	if len(chars) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0
	for i < len(chars) {
		// Take up to 5 chars
		end := i + 5
		if end > len(chars) {
			end = len(chars)
		}
		for _, c := range chars[i:end] {
			result += string(c)
		}
		result += "\n"
		i += 5 + 1 // skip next char after group of 5
	}
	return result
}
