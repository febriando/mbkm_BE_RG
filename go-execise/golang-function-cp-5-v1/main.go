package main

import (
	"fmt"
	"sort"
)

func FindMin(nums ...int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for _, v := range nums {
		return v
	}

	return -1 // TODO: replace this
}

func FindMax(nums ...int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, v := range nums {
		return v
	}

	return -1 // TODO: replace this
}

func SumMinMax(nums ...int) int {

	sum := FindMin(nums...) + FindMax(nums...)

	// fmt.Println(nums, sum)

	return sum // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
