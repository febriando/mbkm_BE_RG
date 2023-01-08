package main

import (
	"fmt"
)

func EmailInfo(email string) string {
	textLength := len(email)
	var temp, tld string

	for i := 0; i < textLength; i++ {
		if string(email[i]) == "@" {
			i++
			for j := i; j < textLength; j++ {
				if string(email[j]) == "." {
					break
				} else {
					temp += string(email[j])
				}

			}
		}
	}

	for k := 0; k < textLength; k++ {
		if string(email[k]) == "." {
			k++
			for l := k; l < textLength; l++ {
				if string(email[l]) == "." {
					tld += "."
				} else {
					tld += string(email[l])
				}

				k = l

			}
		}

	}

	return "Domain: " + temp + " dan TLD: " + tld // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
