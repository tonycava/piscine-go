package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	var temp string
	for index, w := range s {

		if w != ' ' {
			temp = temp + string(w)
		}
		if w == 32 || w == 11 || w == 9 || w == 13 {
			if temp != "" {
				tab = append(tab, temp)
				temp = ""
			}
		}
		if index == len(s)-1 {
			tab = append(tab, temp)
		}
	}
	return tab
}
