package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	var header, position, index, value string
	var makeMap = make(map[string][]string)
	for i := 0; i < len(data); i++ {
		splits := strings.Split(data[i], "-")
		header = splits[0]
		index = splits[1]
		intIndex, _ := strconv.Atoi(index)
		position = splits[2]
		value = splits[3]

		if index == "0" {
			if position == "first" {
				makeMap[header] = append(makeMap[header], value)
			} else {
				makeMap[header][intIndex] = makeMap[header][intIndex] + " " + value
			}
		} else {
			if position == "first" {
				makeMap[header] = append(makeMap[header], value)
			} else {
				makeMap[header][intIndex] = makeMap[header][intIndex] + " " + value
			}
		}
	}
	return (makeMap)
}

// func ChangeOutput(data []string) map[string][]string {
// 	// var makeMap map[string]string
// 	// makeMap := make(map[string]string)
// 	lengthData := len(data)
// 	var splitData []string
// 	// var dataNameFirst []string
// 	// var dataNameLast []string
// 	var dataName [][]string
// 	var newMap = map[string]string{}
// 	var accountMap map[string][]string
// 	// var result map[string][]string
// 	// var newMap map[string][]string
// 	for i := 0; i < lengthData; i++ {
// 		splitData = strings.Split(data[i], "-")
// 		dataName = append(dataName, splitData)
// 		for j := 0; j < len(dataName[i]); j++ {
// 			newMap[dataName[i][0]] = dataName[i][1]
// 			newMap[dataName[i][2]] = dataName[i][3]
// 		}
// 		// result = accountMap
// 		// fmt.Println(newMap["last"])
// 		// for _, v := range splitData {
// 		// 	aMap := map[string][]string{
// 		// 		splitData[0], [splitData[2]: splitData[3
// 		// 	}
// 		// 	newMap[v] = append(newMap[v], aMap)
// 		// }
// 		// fmt.Println(splitData)
// 		fmt.Println(newMap)
// 	}
// 	fmt.Println(dataName)
// 	// fmt.Println(splitData, len(splitData))
// 	return accountMap // TODO: replace this
// }
// func ChangeOutput(data []string) map[string][]string {
// 	// var header, indexs, position, values string
// 	var header, position, values string
// 	// var arrName []string
// 	var temp []string
// 	// var result []string
// 	makeMap := make(map[string][]string)
// 	for i := 0; i < len(data); i++ {
// 		splits := strings.Split(data[i], "-")
// 		header = splits[0]
// 		// indexs = splits[1]
// 		position = splits[2]
// 		values = splits[3]
// 		name := ""
// 		makeMap[header] = append(makeMap[header], values)
// 		fmt.Println(makeMap)
// 		// for j := 0; j < i; j++ {
// 		// 	if header[i] == header[j] && indexs[i] == indexs[j] {
// 		// 		arrName = append(arrName, values[i]+" "+values[j])
// 		// 		fmt.Println(arrName)
// 		// 	}
// 		// }
// 		// if i%2 == 0 {
// 		// 	temp = append(temp, splits...)
// 		// } else {
// 		// 	// makeMap[header] = []string{temp[3], splits[3]}
// 		// 	bla := []string{temp[3], splits[3]}
// 		// 	makeMap[header] = bla
// 		// 	makeMap[header] = append(bla, bla...)
// 		// 	temp = []string{}
// 		// 	fmt.Println(makeMap)
// 		// }
// 	}
// 	return nil
// }

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
