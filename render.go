package main

// returns the full art of the input
func RenderWord(word string, banner []string) string {
	var result string
	if len(word) == 0 {
		return result
	}
	/* The terminal prints top to bottom and can’t go back
	up, so each row has to be completed across all
	 characters before moving on */
	for row := range 8 {
		line := ""
		for _, char := range word {
			line = line + banner[CharStart(char)+row]
		}
		result = result + line + "\n"
	}
	return result
}
