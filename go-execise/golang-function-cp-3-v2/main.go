package main

import (
	"fmt"
	"strings"
)

// var splitWords []string

func FindShortestName(names string) string {

	splitWords := []string{}

	for i := 0; i < len(names); i++ {
		if string(names[i]) == " " {
			splitWords = strings.Split(string(names), " ")

			for j := 0; j < len(splitWords); j++ {

				if len(splitWords) == 2 {
					return splitWords[j+1]
				}

				if len(splitWords[j]) < len(splitWords[j+1]) && (splitWords[j] != splitWords[j+1]) {
					resultSplit := splitWords[j]
					return resultSplit
				}

			}
		}

		if string(names[i]) == ";" {
			splitWords = strings.Split(string(names), ";")
			for j := 0; j < len(splitWords); j++ {
				if len(splitWords[j]) < len(splitWords[j+1]) {
					return splitWords[j]
				} else if len(splitWords[j]) > len(splitWords[j+1]) {
					return splitWords[j+1]
				}
			}

		}

		if string(names[i]) == "," {
			splitWords = strings.Split(string(names), ",")
			for j := 0; j < len(splitWords); j++ {
				if len(splitWords[j]) <= len(splitWords[j+1]) && (splitWords[j] != splitWords[j+1]) {
					return splitWords[j]
				} else {
					return splitWords[j+1]
				}
			}

		}
	}
	return ""

	// TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko"))
	fmt.Println(FindShortestName("Hanif;Joko;Indah;Ari;Intan"))
	fmt.Println(FindShortestName("A,B,C,D,E"))
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))
}
