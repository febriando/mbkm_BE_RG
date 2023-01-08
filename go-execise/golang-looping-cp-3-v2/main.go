package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	text1 := strings.ToUpper(text)
	textLength := len(text)
	hasil := 0

	for i := 0; i < textLength; i++ {
		if string(text1[i]) == "R" || string(text1[i]) == "S" || string(text1[i]) == "T" || string(text1[i]) == "Z" {
			hasil++
		}
	}

	// return "" // TODO: replace this

	return hasil
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Rerursturiru"))
}
