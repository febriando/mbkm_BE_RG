package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {

	// var arrThird []string
	// var mapSplit map[string]string

	for i := 0; i < len(data); i++ {
		splitData := strings.Split(data[i], "-")
		
    mapSplit := make(map[string]string)
    
    // data3 := splitData[3]
		// arrThird = append(arrThird, data3)
		// fmt.Println(arrThird)
	}

	return nil // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
