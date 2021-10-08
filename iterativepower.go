package piscine

func IterativePower(nb int, power int) int {
	back := 1
	if power <= 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		back = back * nb
	}
	return back
}
