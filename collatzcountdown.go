package piscine

func CollatzCountdown(start int) int {
	var steps int

	for start != 1 {
		if start%2 == 0 {
			start /= 2
			steps++
		}
		if start%2 == 1 {
			start *= 3
			start++
			steps++
		}
	}
	return steps
}
