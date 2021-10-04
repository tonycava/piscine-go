package piscine

func UltimateDivMod(a *int, b *int) {
	var c *int
	*c = *a / *b
	*a = *c
	*c = *a % *b
	*b = *c
}
