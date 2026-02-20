package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)

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
		taken := 0
		for i < len(runes) && taken < 5 {
			if runes[i] != ' ' {
				result = append(result, runes[i])
				taken++
			}
			i++
		}
		if taken == 0 {
			break
		}
		if i < len(runes) {
			result = append(result, runes[i])
			i++
		}
		result = append(result, '\n')
	}
	return string(result)
}
