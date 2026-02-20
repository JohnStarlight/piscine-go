package piscine

func CollatzCountdown(start int) int {
	steps := 0

	if start <= 0 {
		return -1
	}
	for start != 1 {
		if start%2 == 0 {
			start /= 2
			steps++
		} else {
			start *= 3
			start++
			steps++
		}
	}
	return steps
}
