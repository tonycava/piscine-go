package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	var temp string
	for index, w := range s {
		if w != ' ' && w != 11 && w != 9 && w != 13 {
			temp = temp + string(w)
		}
		if w == 32 || w == 11 || w == 9 || w == 13 {
			if temp != "" {
				tab = append(tab, temp)
				temp = ""
			}
		}
		if index == len(s)-1 {
			if temp != "" {
				tab = append(tab, temp)
			}
		}
	}
	return tab
}
