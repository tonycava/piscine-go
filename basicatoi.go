package piscine

import (
	"fmt"
	"strconv"
)

func BasicAtoi(s string) int {
	i1, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println(i1)
	}
	return i1
}
