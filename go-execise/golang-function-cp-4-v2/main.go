package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {

	hasil := ""

	// index := contains(data, input)

	for i := 0; i < len(data); i++ {

		resultContains := false

		resultContains = (strings.Contains(data[i], input))

		if resultContains == true && i == len(data)-1 {
			hasil += data[i]
		} else if resultContains == true && i == 0 {
			hasil += data[i]
		} else if resultContains == true {
			hasil += data[i] + ","
		}

	}
	return hasil // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
	fmt.Println(FindSimilarData("mobil", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
}
