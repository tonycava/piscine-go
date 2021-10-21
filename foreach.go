package piscine

func ForEach(f func(int), a []int) {
	test := ""
	for _, salut := range a {
		test = test + string(salut)
	}
}
