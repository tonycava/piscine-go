package piscine

func ConcatParams(args []string) string {
	var tab string
	for i := 0; i < len(args); i++ {
		tab = tab + string(args[i])
		if i < len(args)-1 {
			tab = tab + string("\n")
		}
	}
	return tab
}
