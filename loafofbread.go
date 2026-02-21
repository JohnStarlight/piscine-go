package piscine

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
		// Μάζεψε 5 non-space χαρακτήρες
		count := 0
		for i < len(str) && count < 5 {
			if str[i] != ' ' {
				result += string(str[i])
				count++
			}
			i++
		}
		// Skip 1: αν είναι space, το skip γίνεται φυσικά, αλλιώς skip το γράμμα
		for i < len(str) && str[i] == ' ' {
			i++
		}
		// skip το ένα γράμμα
		if i < len(str) {
			i++
		}
	}
	return result + "\n"
}
