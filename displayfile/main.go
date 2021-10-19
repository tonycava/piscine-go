package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	os.Stdout.Write(data)
}
