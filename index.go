package piscine

func Index(s string, toFind string) int {
	for index := 0; len(toFind)+index <= len(s); index++ {
		if s[index:len(toFind)+index] == toFind {
			return index
		}
	}
	return -1
}
