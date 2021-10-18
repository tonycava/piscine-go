package piscine

func MakeRange(min, max int) []int {
	var tab []int
	if min >= max {
		return nil
	} else {
		tale := make([]int, max-min)
		for i := 0; i < (max - min); i++ {
			tale[i] = min + i
		}
	}
	return tab
}
