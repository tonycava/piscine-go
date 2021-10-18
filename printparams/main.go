package piscine

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args

	for i := 1; i < len(arg); i++ {
		for _, w := range arg[i] {
			z01.PrintRune(w)
		}
	}
	z01.PrintRune('\n')
}
