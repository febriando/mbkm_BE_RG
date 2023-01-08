package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {

	splitWord := strings.Split(*words, "")
	slurAlpha := 0

	for i := 0; i < len(splitWord); i++ {
		if string(splitWord[i]) == "s" {
			slurAlpha++
			*words = strings.Replace(*words, "s", "l", slurAlpha)
		}
		if string(splitWord[i]) == "S" {
			slurAlpha++
			*words = strings.Replace(*words, "S", "L", slurAlpha)
		}
		if string(splitWord[i]) == "r" {
			slurAlpha++
			*words = strings.Replace(*words, "r", "l", slurAlpha)
		}
		if string(splitWord[i]) == "R" {
			slurAlpha++
			*words = strings.Replace(*words, "R", "L", slurAlpha)
		}
		if string(splitWord[i]) == "z" {
			slurAlpha++
			*words = strings.Replace(*words, "z", "l", slurAlpha)
		}
		if string(splitWord[i]) == "Z" {
			slurAlpha++
			*words = strings.Replace(*words, "Z", "L", slurAlpha)
		}
	}

	// return *words // TODO: answer here
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Steven"
	// fmt.Println(&words)
	SlurredTalk(&words)
	fmt.Println(words)
}
