package piscine

func LoafOfBread(str string) string {
	res := ""
	if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		for i := len(str); i < 5; i++ {
			if str[i] == ' ' {
				res = res + string(str[i])
			}
		}
		return res
	}
}
