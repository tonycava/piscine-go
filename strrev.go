package piscine

func StrRev(s string) string {
	runes := []rune(s)

	for i := 0; i < len(s)/2; i++ {
		reserve := runes[i]
		runes[i] = runes[len(runes)-i-1]
		runes[len(runes)-i-1] = reserve
	}
	return string(runes)
}
