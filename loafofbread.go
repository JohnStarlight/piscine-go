package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	ns := ""
	for _, r := range str {
		if r != ' ' {
			ns += string(r)
			if len(ns) == 5 {
				break
			}
		}
	}

	lastSpace := -1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			lastSpace = i
			break
		}
	}

	if lastSpace != -1 {
		return ns + str[lastSpace:] + "\n"
	}
	return ns + "\n"
}
