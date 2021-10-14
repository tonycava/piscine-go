package piscine

func SortIntegerTable(table []int) {
	var first int
	var second int

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-1; j++ {
			first = table[j]
			second = table[j+1]
			if first > second {
				table[j] = second
				table[j+1] = first
			}
		}
	}
}
