package piscine

func LoafOfBread(str string) string {
	// count non-space
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
	collected := 0

	for i := 0; i < len(str); i++ {
		// copy char always
		result += string(str[i])

		if str[i] != ' ' {
			collected++
		}

		// when 5 reached → skip next char
		if collected == 5 {
			collected = 0
			i++ // skip
		}
	}

	return result + "\n"
}
