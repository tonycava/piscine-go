package piscine

func AppendRange(min, max int) []int {
	var tab []int
	if min > max {
		return tab
	} else {
		for _, nub := range tab {
			tab = append(tab, nub+1)
		}
	}
	return tab
}
