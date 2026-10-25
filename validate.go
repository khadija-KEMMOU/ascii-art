package main

import "fmt"

// returns error if the input had a rune in it
func ValidateInput(s string) error {
	for _, char := range s {
		if char < 32 || char > 126 {
			//the %q prints a rune in quotes
			return fmt.Errorf("unsupported character %q", char)
		}
	}
	return nil
}
