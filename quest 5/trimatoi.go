package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			sign = -1
		} else if r >= '0' && r <= '9' {
			foundDigit = true
			digit := int(r - '0')
			result = result*10 + digit
		}
	}

	if !foundDigit {
		return 0
	}

	return result * sign
}
