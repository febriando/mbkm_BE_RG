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

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func GetDayDifference(date string) int {

	var splitDate []string
	monthMap := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"Agust":     8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
	var jarakHari int

	// hari := 1

	for i := 0; i < len(date); i++ {
		splitDate = strings.Split(date, " ")
		year, _ := strconv.Atoi(string(splitDate[5]))
		tgl, _ := strconv.Atoi(string(splitDate[0]))
		tgl2, _ := strconv.Atoi(string(splitDate[3]))

		// tglFix := (tgl2 - tgl) + 1

		t1 := Date(year, monthMap[string(splitDate[1])], tgl)
		t2 := Date(year, monthMap[string(splitDate[4])], tgl2)

		days := t2.Sub(t1).Hours() / 24
		jarakHari = int(days)

		if jarakHari <= 1 {
			jarakHari += 1
		} else if jarakHari >= 1 {
			jarakHari += 1
		}
	}

	return jarakHari // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {

	var nameMap = map[string]int{}

	for i := 0; i < rangeDay; i++ {
		for j := range data[i] {
			nameMap[data[i][j]]++
		}
	}

	var karyawanDanGajiHarian = map[string]string{}

	for k, v := range nameMap {
		if v > 0 {
			v = v * 50000
		}
		karyawanDanGajiHarian[k] = FormatRupiah(v)
	}

	return karyawanDanGajiHarian // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {

	rangeDay := GetDayDifference(dateRange)
	result := GetSalary(rangeDay, data)

	return result // TODO: replace this
}

func main() {
	res := GetSalaryOverview("21 February - 21 February 2021", [][]string{
		{"A", "B"},
		{"A", "C"},
		{"A", "D"},
	})

	fmt.Println(res)

	fmt.Println(GetDayDifference("30 March - 2 May 2021"))

}
