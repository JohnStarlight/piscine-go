package piscine

func Capitalize(s string) string {
	r := []rune(s)
	isNewWord := true

	for i := 0; i < len(r); i++ {
		if isAlphaNum(r[i]) {
			if isNewWord {
				// Capitalize
				if r[i] >= 'a' && r[i] <= 'z' {
					r[i] = r[i] - 32
				}
				isNewWord = false
			} else {
				// Lowercase
				if r[i] >= 'A' && r[i] <= 'Z' {
					r[i] = r[i] + 32
				}
			}
		} else {
			isNewWord = true
		}
	}

	return string(r)
}

func isAlphaNum(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}
