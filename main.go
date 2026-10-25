package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage : go run . \"text\"")
		os.Exit(1)
	}
	input := os.Args[1]
	if len(input) == 0 {
		return
	}
	err := ValidateInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	standard, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pieces := strings.Split(input, `\n`)
	for _, piece := range pieces {
		if len(piece) == 0 {
			fmt.Println()
		} else {
			fmt.Print(RenderWord(piece, standard))
		}
	}
}
