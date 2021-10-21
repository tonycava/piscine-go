package piscine

func Map(f func(int) bool, a []int) []bool {
	var tab []bool
	for _, Tprime := range a {
		tab = append(tab, f(Tprime))
	}
	return tab
}
