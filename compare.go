package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return Compare(a, b)
}