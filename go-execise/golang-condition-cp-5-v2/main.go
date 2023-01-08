package main

import "fmt"

func TicketPlayground(height, age int) int {

	if age < 5 {
		return -1
	} else if age > 12 {
		return 100000
	} else if age == 12 || height >= 160 {
		return 60000
	} else if age >= 10 || height > 150 {
		return 40000
	} else if age >= 8 || height > 135 {
		return 25000
	} else if age >= 5 || height > 120 {
		return 15000
	}

	return 0 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(85, 4))
	fmt.Println(TicketPlayground(140, 13))
	fmt.Println(TicketPlayground(121, 7))
	fmt.Println(TicketPlayground(136, 7))
	fmt.Println(TicketPlayground(151, 6))
	fmt.Println(TicketPlayground(161, 7))
	fmt.Println(TicketPlayground(121, 9))
	fmt.Println(TicketPlayground(136, 9))
	fmt.Println(TicketPlayground(151, 9))
	fmt.Println(TicketPlayground(161, 9))
	fmt.Println(TicketPlayground(121, 11))
	fmt.Println(TicketPlayground(136, 11))
	fmt.Println(TicketPlayground(151, 11))
	fmt.Println(TicketPlayground(161, 11))
	fmt.Println(TicketPlayground(121, 12))
	fmt.Println(TicketPlayground(136, 12))
	fmt.Println(TicketPlayground(151, 12))
	fmt.Println(TicketPlayground(161, 12))
}
