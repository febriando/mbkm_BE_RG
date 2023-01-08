package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {

	sort.Slice(height[:], func(i, j int) bool {
		return height[i] < height[j]
	})

	return height // TODO: replace this
}

func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159}))
}
