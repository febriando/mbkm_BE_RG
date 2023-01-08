package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	harga := (VIP * 30) + (regular * 20) + (student * 10)
	// harga1 := float64(harga)
	tiket := VIP + regular + student
	day = day % 2

	if harga <= 100 {
		return float32(harga)
	} else {
		for day == 1 {
			if tiket < 5 {
				harga = harga - (harga * 15 / 100)
				harga2 := int64(harga)
				return float32(harga2)
			} else if tiket >= 5 {
				harga = harga - (harga * 25 / 100)
				harga2 := int64(harga)
				return float32(harga2)
			}
		}

		for day == 0 {
			if tiket < 5 {
				harga = harga - (harga * 10 / 100)
				harga2 := int64(harga)
				return float32(harga2)
			} else if tiket >= 5 {
				harga = harga - (harga * 20 / 100)
				harga2 := int64(harga)
				return float32(harga2)
			} else {
				return 111
			}
		}
	}

	// if condition {

	// }

	return 0 // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(4, 0, 0, 21))
	fmt.Println(GetTicketPrice(0, 0, 5, 1))
	fmt.Println(GetTicketPrice(1, 0, 6, 15))
	fmt.Println(GetTicketPrice(4, 0, 0, 13))
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
	fmt.Println(GetTicketPrice(4, 4, 4, 20))
	fmt.Println(GetTicketPrice(4, 0, 0, 20))
	fmt.Println(GetTicketPrice(3, 3, 3, 22))
	fmt.Println(GetTicketPrice(4, 4, 4, 22))
}
