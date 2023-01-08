package main

import (
	"fmt"
)

func CountProfit(data [][][2]int) []int {

	empty := make([]int, 0, len(data))
	if len(data) == 0 {
		return empty
	}

	var lengthI int
	var lengthIJ int
	var nilaiK int
	arrBulanAll := []int{}

	for i := range data {
		lengthI = len(data[i])
		for j := range data[i] {
			lengthIJ = len(data[i][j])
			for k := range data[i][j] {

				if k > 0 {
					nilaiK = data[i][j][0] - data[i][j][1]
					arrBulanAll = append(arrBulanAll, nilaiK)
				}
				// fmt.Println(lengthI, lengthIJ, -1)
			}
		}
	}

	result := 0
	if lengthI < lengthIJ {
		for _, v := range arrBulanAll {
			result += v
		}
	}
	if lengthI > lengthIJ {
		return arrBulanAll
	}
	if lengthI == lengthIJ {
		for _, v := range arrBulanAll {
			result += v
		}
	}

	hasilArr := []int{}
	hasilArr = append(hasilArr, result)

	return hasilArr // TODO: replace this
}

func main() {
	// data := [][][2]int{{{1000, 800}, {700, 500}}, {{1000, 800}, {900, 200}}}
	// data := [][][2]int{{{1000, 200}}, {{500, 100}}, {{600, 100}}, {{450, 150}}, {{100, 50}}}
	data := [][][2]int{}

	fmt.Println(CountProfit(data))
}
