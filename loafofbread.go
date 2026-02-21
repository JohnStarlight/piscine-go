func LoafOfBread(str string) string {
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
		// Skip: αν είναι space, αυτό μετράει ως skip
		// αν είναι γράμμα, skip αυτό
		if i < len(str) {
			i++
		}
	}
	return result + "\n"
}