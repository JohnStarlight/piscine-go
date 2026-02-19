package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	word := ""
	counts := make(map[string]int)

	for _, v := range str {
		if v != ' ' {
			word += string(v)
		} else {
			if word != "" {
				counts[word]++
			}
			word = ""
		}
	}

	if word != "" {
		counts[word]++
	}

	return counts
}
