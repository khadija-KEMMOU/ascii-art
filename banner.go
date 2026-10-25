package main

import (
	"errors"
	"os"
	"strings"
)

// reads the banner file and converts it to a slice of strings with no newlines
func LoadBanner(path string) ([]string, error) {
	banner, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	standard := string(banner)
	lines := strings.Split(standard, "\n")
	// checks the file has at least enough lines for every character
	/* it doesnt check if the file is right, this check is
	only to avoid the index out of range crash */
	if len(lines) < 855 {
		return nil, errors.New("banner file is too short")
	}
	return lines, nil
}

// looks for the start of the shape of a character in the banner
func CharStart(c rune) int {
	rowStart := (c-32)*9 + 1
	return int(rowStart)
}
