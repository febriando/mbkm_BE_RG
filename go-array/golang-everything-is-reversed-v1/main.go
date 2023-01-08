package main

import (
	"fmt"
	"sort"
)

func ReverseData(arr [5]int) [5]int {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]

		newArr := arr
		for n := 0; n < len(newArr); n++ {
			sort.Slice(newArr[:], func(k, l int) bool {
				return newArr[k] > newArr[l]
			})
			for _, v := range arr {
				fmt.Println(v)
			}
		}

	}

	return arr // TODO: replace this
}

func main() {
	arrTests := [5]int{123, 456, 11, 1, 2}
	fmt.Println(ReverseData(arrTests))
}
