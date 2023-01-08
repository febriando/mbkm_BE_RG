package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	strLength := len(str)
	var str1, hasil string
	str2 := strings.ToLower(str)
	for i := strLength - 1; i >= 0; i-- {
		str1 += string(str2[i])
		// fmt.Println(str1)
	}

	strDaristr1 := []rune(str)
	str3 := []rune(str1)
	var changeLowerCase rune

	for i := 0; i < strLength; i++ {
		if unicode.IsUpper(strDaristr1[i]) {
			changeLowerCase += unicode.ToUpper(rune(str3[i]))
			hasil += string(changeLowerCase)
		} else {
			hasil += string(strDaristr1)
		}
	}
	// return the reversed string.
	return string(hasil)

	// return "random" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Print(ReverseWord("Aku Sayang Ibu"))
	// fmt.Println(ReverseWord("A bird fly to the Sky"))
}
