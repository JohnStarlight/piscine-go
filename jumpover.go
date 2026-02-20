package piscine

func JumpOver(str string) string {
	result := ""
	if len(str) < 2 || len(str) == 0 {
		result = "\n"
		return result
	}
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	result += "\n"
	return result
}
