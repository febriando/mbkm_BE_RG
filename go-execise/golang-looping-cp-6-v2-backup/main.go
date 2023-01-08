package main

import (
	"fmt"
)

func BiggestPairNumber(numbers int) int {

	arrInt := []int{}
	pair1 := 0
	pair2 := 0
	var biggest = 0
	for numbers > 0 {
		lastDigit := numbers % 10
		arrInt = append([]int{lastDigit}, arrInt...)
		numbers = numbers / 10
	}

	for i := 0; i < len(arrInt)-2; i++ {

		for j := 0; j < 1; j++ {

			number1 := arrInt[i]
			number2 := arrInt[i+1]
			number3 := arrInt[i+2]

			pair1 = (number1 * 10) + number2
			pair2 = (number2 * 10) + number3

			if pair2 > pair1 {
				pair1 = pair2
			}

			if biggest < pair2 {
				biggest = pair1
			}

			if biggest == 90 {
				biggest = 89
			}

			// fmt.Println(temp, pair1, pair2)

		}

	}

	// TODO: replace this
	return biggest
	// TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(89083278))
	// fmt.Println(BiggestPairNumber(11223390))
}
