package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	word := ""

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	if len(words) == 0 {
		return nil
	}
	return words
}
