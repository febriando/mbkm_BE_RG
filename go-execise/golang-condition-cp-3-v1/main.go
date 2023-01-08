package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {

	nilai := (math + science + english + indonesia) / 4

	for i := 0; i <= 100; {
		if nilai == 100 {
			return "Sempurna"
		} else if nilai >= 90 {
			return "Sangat Baik"
		} else if nilai >= 80 {
			return "Baik"
		} else if nilai >= 70 {
			return "Cukup"
		} else if nilai >= 60 {
			return "Kurang"
		} else {
			return "Sangat kurang"
		}

	}

	return "" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(100, 100, 100, 100))
	fmt.Println(GetPredicate(90, 90, 90, 90))
	fmt.Println(GetPredicate(92, 92, 95, 100))
	fmt.Println(GetPredicate(80, 80, 80, 80))
	fmt.Println(GetPredicate(82, 85, 95, 80))
	fmt.Println(GetPredicate(75, 75, 75, 75))
	fmt.Println(GetPredicate(70, 70, 70, 90))
	fmt.Println(GetPredicate(60, 65, 65, 65))
	fmt.Println(GetPredicate(60, 60, 60, 75))
	fmt.Println(GetPredicate(50, 50, 50, 50))
	fmt.Println(GetPredicate(50, 50, 50, 60))
}
