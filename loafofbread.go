package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}

	nonSpace := 0
	for _, c := range str {
		if c != ' ' {
			nonSpace++
		}
	}
	if nonSpace < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0
	for i < len(str) {
		if result != "" {
			result += " "
		}
		count := 0
		for i < len(str) && count < 5 {
			if str[i] != ' ' {
				result += string(str[i])
				count++
			}
			i++
		}
		if i < len(str) {
			i++
		}
	}

	// Αφαίρεσε trailing spaces
	for len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
