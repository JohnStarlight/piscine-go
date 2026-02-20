package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)
	if len(runes) < 5 {
		return "Invalid Output\n"
	}
	var res []rune
	i := 0
	for i < len(runes) {
		count := 0
		for i < len(runes) && count < 5 {
			if runes[i] != ' ' {
				res = append(res, runes[i])
				count++
			}
			i++
		}
		if count == 0 {
			break
		}
		for i < len(runes) {
			if runes[i] != ' ' {
				res = append(res, ' ')
				i++
				break
			}
			i++
		}
	}
	return string(res) + "\n"
}
