package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {

	word := ""
	result := []map[string]any{}
	for _, v := range data {
		newV := strings.Split(v, ";")
		dataMap := make(map[string]any)

		name := newV[0]
		age, _ := strconv.Atoi(newV[1])
		address := newV[2]
		height, _ := strconv.ParseFloat(newV[3], 64)
		isMarried, _ := strconv.ParseBool(newV[4])
		dataMap["name"] = name
		dataMap["age"] = age
		dataMap["address"] = address
		if newV[3] != "" {
			dataMap["height"] = height
		}
		if newV[4] != "" {
			dataMap["isMarried"] = isMarried
		}
		result = append(result, dataMap)
	}
	fmt.Println(word)

	return result // TODO: replace this
}
func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	fmt.Println(PopulationData(data))
}
