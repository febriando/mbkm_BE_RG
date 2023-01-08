package main

import (
	"fmt"
)

func BiggestPairNumber(numbers int) int {

	arrInt := []int{}
	pair1 := 0
	pair2 := 0
	first := 0
	second := 0
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

			if pair1 > pair2 {
				first = pair1
				second = pair2
			} else {
				first = pair2
				second = pair1
			}

			if pair1 > first {
				second = first
				first = pair1
			} else if pair1 > second && pair1 != first {
				second = pair1
			}

		}

	}

	// TODO: replace this
	return first
	// TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
