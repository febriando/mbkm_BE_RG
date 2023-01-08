package main

import (
	"fmt"
	"math"
	"strconv"
)

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {

	if amount == 0 {
		return []int{}
	}

	hargaTotal := 0
	hargaPrice := 0
	hargaTax := 0

	for i := 0; i < len(products); i++ {
		hargaPrice += products[i].Price
		hargaTax += products[i].Tax
		hargaTotal = hargaPrice + hargaTax
	}

	if amount == hargaTotal {
		return []int{}
	}

	uangKembalianAmount := amount - hargaTotal

	var myArray []int

	//10000 - 8800 = 1200 => 1000 200

	strUangKembalian := strconv.Itoa(uangKembalianAmount)

	// splitHargaTotal := strings.Split(strHargaTotal, "")
	lengthHargaTotal := len(strUangKembalian)
	// fmt.Println(lengthHargaTotal, strHargaTotal, splitHargaTotal)

	for j := 0; j < lengthHargaTotal; j++ {

		intUangKembalian := strUangKembalian[j]
		byteToInt, _ := strconv.Atoi(string(intUangKembalian))
		j++

		pangkatUang := byteToInt * int(math.Pow10(lengthHargaTotal-j))

		if pangkatUang == 0 {
			continue
		} else if pangkatUang == 1000 {
			myArray = append(myArray)

		} else if pangkatUang > 500 {
			temp := pangkatUang - 500
			myArray = append(myArray, temp)
			if temp > 200 {
				temp := temp - 200
				myArray = append(myArray, temp)
			}
		}

		j--
	}

	// result := []int{intUangKembalian}

	return myArray // TODO: replace this
}

func main() {
	fmt.Println(MoneyChanges(30000, []Product{
		{Name: "baju 1", Price: 10000, Tax: 1000}, {Name: "Sepatu", Price: 15550, Tax: 1555},
	}))

}
