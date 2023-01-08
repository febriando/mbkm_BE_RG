package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(date1 []int, date2 []int) []int {

	newData := append(date1, date2...)

	sort.Slice(newData[:], func(k, l int) bool {
		return newData[k] < newData[l]
	})

	sameDate := []int{}
	for i := 0; i < len(newData)-1; i++ {
		left := newData[i]
		right := newData[i+1]
		if left == right {
			sameDate = append(sameDate, right)
			// fmt.Println(sameDate)
		}
	}

	return sameDate // TODO: replace this
}

func main() {
	Imam := []int{11, 12, 13, 14, 15, 22}
	Permana := []int{5, 10, 12, 13, 20, 21, 22}
	fmt.Println(SchedulableDays(Imam, Permana))
}
