package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {

	var strDay, strMonth, strYear string

	if day <= 9 {
		strDay = strconv.Itoa(day)
		strDay = "0" + strDay
	} else {
		strDay = strconv.Itoa(day)
	}

	if month == 1 {
		strMonth = strconv.Itoa(month)
		strMonth = "January"

	} else if month == 2 {
		strMonth = strconv.Itoa(month)
		strMonth = "February"

	} else if month == 3 {
		strMonth = strconv.Itoa(month)
		strMonth = "March"

	} else if month == 4 {
		strMonth = strconv.Itoa(month)
		strMonth = "April"

	} else if month == 5 {
		strMonth = strconv.Itoa(month)
		strMonth = "May"

	} else if month == 6 {
		strMonth = strconv.Itoa(month)
		strMonth = "June"

	} else if month == 7 {
		strMonth = strconv.Itoa(month)
		strMonth = "July"

	} else if month == 8 {
		strMonth = strconv.Itoa(month)
		strMonth = "August"

	} else if month == 9 {
		strMonth = strconv.Itoa(month)
		strMonth = "September"

	} else if month == 10 {
		strMonth = strconv.Itoa(month)
		strMonth = "October"

	} else if month == 11 {
		strMonth = strconv.Itoa(month)
		strMonth = "November"

	} else if month == 12 {
		strMonth = strconv.Itoa(month)
		strMonth = "December"

	}

	if year > 0 {
		strYear = strconv.Itoa(year)
	}

	return string(strDay + "-" + strMonth + "-" + strYear) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(10, 1, 2012))
}
