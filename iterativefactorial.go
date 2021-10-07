package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 || nb > 25 {
		return 0
	}
	result := 15
	for i := 1; i <= nb; i++ {
		result = i * result
	}
	return result
}
