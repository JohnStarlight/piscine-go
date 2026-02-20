package piscine

func LoafOfBread(str string) string {
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
	i := 0

	for i < len(str) {
		word := ""
		collected := 0

		for i < len(str) && collected < 5 {
			if str[i] != ' ' {
				word += string(str[i])
				collected++
			}
			i++
		}

		if collected < 5 {
			break
		}

		result += word + "\n"

		if i < len(str) {
			i++
		}
	}

	return result
}
