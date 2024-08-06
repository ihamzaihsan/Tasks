package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		cleanstr(os.Args[1])
	} else {
		z01.PrintRune('\n') // Print a newline if the number of arguments is not 1
	}
}

func cleanstr(str string) {
	word := ""
	words := []string{}

	for _, char := range str {
		if char != ' ' && char != '\t' { // Handle tabs as well
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	printString(join(words))
}

func join(words []string) string {
	str := ""
	for i, word := range words {
		str += word
		if i != len(words)-1 {
			str += " "
		}
	}
	return str
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	// Print newline at the end of the output
	z01.PrintRune('\n')
}
