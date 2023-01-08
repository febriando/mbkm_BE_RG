package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {

	vokal := 0
	nonVokal := 0
	var checkerVokal bool
	// writeStr := strings.Split()

	// len(str)

	for i := 0; i < len(str); i++ {

		if strings.ToLower(string(str[i])) == "a" || strings.ToLower(string(str[i])) == "e" || strings.ToLower(string(str[i])) == "i" || strings.ToLower(string(str[i])) == "o" || strings.ToLower(string(str[i])) == "u" {
			vokal += 1
		} else if str[i] >= 32 && str[i] <= 64 {
			nonVokal = nonVokal
		} else {
			nonVokal += 1
		}

		if vokal == 0 {
			checkerVokal = true
		} else {
			checkerVokal = false
		}

	}

	return vokal, nonVokal, checkerVokal // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("HALO halo"))
}
