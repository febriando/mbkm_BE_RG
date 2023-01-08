package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi ini akan mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	var result string

	numStr := strconv.Itoa(number)

	for i := len(numStr) - 1; i >= 0; i-- {
		result = string(numStr[i]) + result
		if (len(numStr)-i)%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}

func GetDayDifference(date string) int {

	dateByDay := strings.Split(date, " ")
	dayStart := dateByDay[0]
	dayEnd := dateByDay[3]
	monthStart := dateByDay[1]
	monthEnd := dateByDay[4]
	year := dateByDay[5]
	dateByDayStart := time.Date(year, 2, 1, 12, 30, 0, 0, time.UTC)
	dateByDayEnd := time.Date(2000, 2, 3, 12, 30, 0, 0, time.UTC)
	year1, month1, day1 := dateByDayStart.Date()
	year2, month2, day2 := dateByDayEnd.Date()

	fmt.Println((day2 - day1) + 1)
	fmt.Println(year1, month1, day1)
	fmt.Println(year2, month2, day2)

	// intDayStart := strconv.Atoi(dayStart)
	// strconv.Atoi(dayEnd)
	// fmt.Println(dayStart, dayEnd, monthStart, monthEnd, year)
	// fmt.Println((dayEnd - dayStart) + 1)

	return 0 // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	return nil // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	return nil // TODO: replace this
}

func main() {
	newRes := "21 February - 23 February 2021"
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	GetDayDifference(newRes)

	fmt.Println(res)
}