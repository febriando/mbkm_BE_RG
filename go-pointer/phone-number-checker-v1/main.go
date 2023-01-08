package main

import (
	"fmt"
)

// var numberToInt int

func PhoneNumberChecker(number string, result *string) {

	var nomorKamu string

	// noTelkomsel := map[string]string{
	// 	"62811": "Telkomsel",
	// 	"62812": "Telkomsel",
	// 	"62813": "Telkomsel",
	// 	"62814": "Telkomsel",
	// 	"62815": "Telkomsel",
	// }

	if len(number) <= 9 || len(number) >= 13 {
		nomorKamu = "invalid"
	}

	// for i := 0; i < len(number); i++ {
	// 	if number[0:5] == noTelkomsel["62811"] {
	// 		nomorKamu = "Telkomsel"
	// 	}
	// }

	switch string(number[0:4]) {
	case "0811", "0812", "0813", "0814", "0815":
		nomorKamu = "Telkomsel"
	case "0816", "0817", "0818", "0819":
		nomorKamu = "Indosat"
	case "0821", "0822", "0823":
		nomorKamu = "XL"
	case "0827", "0828", "0829":
		nomorKamu = "Tri"
	case "0852", "0853":
		nomorKamu = "AS"
	case "0881", "0882", "0883", "0884", "0885", "0886", "0887", "0888":
		nomorKamu = "Smartfren"
	}

	switch string(number[0:5]) {
	case "62811", "62812", "62813", "62814", "62815":
		nomorKamu = "Telkomsel"
	case "62816", "62817", "62818", "62819":
		nomorKamu = "Indosat"
	case "62821", "62822", "62823":
		nomorKamu = "XL"
	case "62827", "62828", "62829":
		nomorKamu = "Tri"
	case "62852", "62853":
		nomorKamu = "AS"
	case "62881", "62882", "62883", "62884", "62885", "62886", "62887", "62888":
		nomorKamu = "Smartfren"
	}

	*result = nomorKamu

	// TODO: answer here
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "62811081634"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
