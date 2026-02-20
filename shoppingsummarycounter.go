package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	word := ""
	counts := make(map[string]int)

	for _, v := range str {
		if v != ' ' {
			word += string(v)
		} else {
			// Remove the 'if word != ""' guard.
			// This allows the map to count "" when it hits multiple spaces.
			counts[word]++
			word = ""
		}
	}

	// Always add the last word, even if it's empty
	counts[word]++

	return counts
}
