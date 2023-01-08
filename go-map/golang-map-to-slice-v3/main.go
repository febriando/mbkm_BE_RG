package main

import (
	"fmt"
)

func MapToSlice(mapData map[string]string) [][]string {
	result := make([][]string, 0, len(mapData))
	for key, value := range mapData {
		result = append(result, []string{key, value})
	}
	return result
}

func main() {
	data := map[string]string{
		"foo": "33",
		"bar": "44",
	}
	fmt.Println(MapToSlice(data))
}
