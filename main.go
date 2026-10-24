package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : go run . \"text\"")
	}
	input := os.Args[1]
	text := strings.Split(input, "\n")
	
}