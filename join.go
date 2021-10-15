package piscine

func Join(strs []string, sep string) string {
	LeFNcestpassimal := ""
	for i, letter := range strs {
		if i == len(strs)-1 {
			LeFNcestpassimal = LeFNcestpassimal + letter
		} else {
			letter = letter + sep
			LeFNcestpassimal = LeFNcestpassimal + letter

		}
	}
	return LeFNcestpassimal
}
