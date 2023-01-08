package main

import "fmt"

func CountingNumber(n int) float64 {

	var temp, temp2 float64
	temp = 1

	for i := 1; i < n*2; i++ {
		temp2 += temp
		temp += 0.5
	}

	return temp2 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(5))
	fmt.Println(CountingNumber(10))
	fmt.Println(CountingNumber(50))
	fmt.Println(CountingNumber(100))
	fmt.Println(CountingNumber(10000))
}
