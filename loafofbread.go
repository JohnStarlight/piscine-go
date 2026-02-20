package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)

	nonSpace := 0
	for _, c := range runes {
		if c != ' ' {
			nonSpace++
		}
	}
	if nonSpace < 5 {
		return "Invalid Output\n"
	}

	result := []rune{}
	i := 0
	first := true
	for i < len(runes) {
		if !first {
			result = append(result, ' ')
		}
		first = false
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
		// skip 1 non-space char (advance past spaces to find it)
		for i < len(runes) {
			if runes[i] != ' ' {
				i++
				break
			}
			i++
		}
	}
	result = append(result, '\n')
	return string(result)
}
