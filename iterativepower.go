package piscine

func IterativePower(nb int, power int) int {
	back := 1
	for i := 1; i <= power; i++ {
		back = back * nb
	}
	return back
}
