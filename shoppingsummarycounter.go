package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	word := ""
	counts := make(map[string]int)

	for _, v := range str {
		if v != ' ' {
			word += string(v)
		} else {
			// This block skips adding anything to the map if word is ""
			if word != "" {
				counts[word]++
			}
			word = ""
		}
	}

	// Final check for the last word
	if word != "" {
		counts[word]++
	}

	return counts
}
