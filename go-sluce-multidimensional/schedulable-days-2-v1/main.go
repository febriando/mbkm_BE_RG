package main

import "fmt"

func SchedulableDays(villager [][]int) []int {

	var freqMap = map[int]int{}
	var res = []int{}
	for i := range villager {
		for j := range villager[i] {
			freqMap[villager[i][j]]++
		}
	}

	for k, v := range freqMap {
		if v == len(villager) {
			res = append(res, k)
		}
	}

	return res // TODO: replace this
}

func main() {
	data := [][]int{
		{7, 12, 19, 22},
		{12, 19, 21, 23},
		{7, 12, 19},
		{12, 19},
	}
	fmt.Println(SchedulableDays(data))
}
