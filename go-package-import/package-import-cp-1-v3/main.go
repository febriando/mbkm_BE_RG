package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {

	token := strings.Split(calculate, " ")
	// fmt.Printf("%#v\n", token)
	if len(calculate) == 0 {
		return 0
	}

	number, err := strconv.ParseFloat(token[0], 32)
	if err != nil {
		return 0
	}

	calc := internal.NewCalculator(float32(number))
	for i := 1; i < len(token); i += 2 {
		operator := token[i]
		operan, err := strconv.ParseFloat(token[i+1], 32)
		if err != nil {
			return 0
		}

		switch operator {
		case "+":
			calc.Add(float32(operan))
		case "-":
			calc.Subtract(float32(operan))
		case "*":
			calc.Multiply(float32(operan))
		case "/":
			calc.Divide(float32(operan))
		default:
			return 0
		}

	}

	return calc.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
