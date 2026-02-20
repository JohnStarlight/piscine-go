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
	collected := 0

	for i := 0; i < len(str); i++ {
		result += string(str[i])

		if str[i] != ' ' {
			collected++
		}

		if collected == 5 {
			collected = 0
			i++
		}
	}

	return result + "\n"
}
