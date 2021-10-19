package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return
	}
	os.Stdout.Write(data)
}
