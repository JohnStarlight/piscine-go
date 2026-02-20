package piscine

// test
func LoafOfBread(str string) string {
	var chars []rune
	for _, c := range str {
		if c != ' ' {
			chars = append(chars, c)
		}
	}

	if len(chars) < 5 {
		return "Invalid Output\n"
	}

	result := []rune{}
	i := 0
	for i < len(chars) {
		if i > 0 {
			result = append(result, ' ')
		}
		end := i + 5
		if end > len(chars) {
			end = len(chars)
		}
		result = append(result, chars[i:end]...)
		i += 5
	}
	result = append(result, '\n')
	return string(result)
}
