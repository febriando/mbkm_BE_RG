package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Reverse(str string) string {

	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}

	return reverseStr // TODO: replace this
}

func Generate(str string) string {
	str = Reverse(str)
	splitStr := strings.Split(str, "")
	// fmt.Println(splitStr)
	vokal := 0
	replaceStr := ""

	for i := 0; i < len(splitStr); i++ {
		switch splitStr[i] {
		case "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z":
			replaceStr += strings.ToUpper(splitStr[i])
		case "B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z":
			replaceStr += strings.ToLower(splitStr[i])
		case "!", "#", "$", "%", "&", "‘", "(", ")", "*", "+", ",", "-", ".", "/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "<", "=", ">", "?", "@":
			replaceStr += splitStr[i]
		}

		if string(splitStr[i]) == "a" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "a", "E", vokal)
			continue
		}
		if string(splitStr[i]) == "e" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "e", "I", vokal)
			continue
		}
		if string(splitStr[i]) == "i" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "i", "O", vokal)
			continue
		}
		if string(splitStr[i]) == "o" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "o", "U", vokal)
			continue
		}
		if string(splitStr[i]) == "u" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "u", "A", vokal)
			continue
		}
		if string(splitStr[i]) == "A" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "A", "e", vokal)
			continue
		}
		if string(splitStr[i]) == "E" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "E", "i", vokal)
			continue
		}
		if string(splitStr[i]) == "I" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "I", "o", vokal)
			continue
		}
		if string(splitStr[i]) == "O" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "O", "u", vokal)
			continue
		}
		if string(splitStr[i]) == "U" {
			vokal++
			replaceStr += strings.Replace(splitStr[i], "U", "a", vokal)
			continue
		}
	}

	return replaceStr // TODO: replace this
}

func CheckPassword(str string) string {

	splitStr := strings.Split(str, "")

	// check

	// fmt.Println(splitStr)
	// var runeStr []rune(splitStr)

	// for i := 0; i < len(splitStr); i++ {
	// 	if len(splitStr) < 7 {
	// 		return "sangat lemah"
	// 	}
	// 	switch len(splitStr) >= 7 {
	// 	case unicode.IsLetter(splitStr) == true && unicode.IsNumber(splitStr[i]) == true:
	// 		return "lemah"
	// 	}
	// }

	for _, checkStr := range str {
		if len(splitStr) < 7 {
			return "sangat lemah"
		}
		switch len(splitStr) >= 7 {
		case unicode.IsLetter(checkStr) == true && unicode.IsNumber(checkStr) == true:
			return "lemah"
		case unicode.IsLetter(checkStr) == true && unicode.IsNumber(checkStr) == true && unicode.IsSymbol(checkStr) == true:
			return "sedang"
		case unicode.IsLetter(checkStr) == true && unicode.IsNumber(checkStr) == true && unicode.IsSymbol(checkStr) == true && len(splitStr) >= 14:
			return "sedang"
		}
	}

	return "" // TODO: replace this
}

func PasswordGenerator(base string) (string, string) {
	passwordKamu := Generate(base)
	levelPassword := CheckPassword(base)
	return passwordKamu, levelPassword // TODO: replace this
}

func main() {

	data := "Semangat Feb" // bisa digunakan untuk melakukan debug
	CheckPassword(data)

	res, check := PasswordGenerator(data)

	fmt.Print(res+", ", check)
}
