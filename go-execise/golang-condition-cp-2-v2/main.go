package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	height2 := float64(height)

	for i := 130; i <= 250; i++ {
		if gender == "laki-laki" {
			return ((height2 - 100) - ((height2 - 100) * 0.1))
		} else if gender == "perempuan" {
			return ((height2 - 100) - ((height2 - 100) * 0.15))
		}
	}

	return 0.0 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("laki-laki", 160))
	fmt.Println(BMICalculator("laki-laki", 155))
	fmt.Println(BMICalculator("perempuan", 165))
	fmt.Println(BMICalculator("perempuan", 170))
	fmt.Println(BMICalculator("perempuan", 160))
	fmt.Println(BMICalculator("perempuan", 155))
}
