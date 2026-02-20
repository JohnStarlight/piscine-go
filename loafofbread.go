package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)

	// count non-space chars
	count := 0
	for _, c := range runes {
		if c != ' ' {
			count++
		}
	}
	if count < 5 {
		return "Invalid Output\n"
	}

	result := []rune{}
	i := 0
	for i < len(runes) {
		// collect 5 non-space chars
		taken := 0
		for i < len(runes) && taken < 5 {
			if runes[i] != ' ' {
				result = append(result, runes[i])
				taken++
			}
			i++
		}
		if taken > 0 {
			result = append(result, '\n')
		}
		// skip 1 non-space char
		for i < len(runes) {
			if runes[i] != ' ' {
				i++
				break
			}
			i++
		}
	}
	return string(result)
}
