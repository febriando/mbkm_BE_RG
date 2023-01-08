package invoice

import (
	"fmt"
	"strconv"
	"strings"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if len(wi.Date) == 0 {
		return InvoiceData{}, fmt.Errorf("invoice date is empty")
	}
	tanggalInvoice := strings.Split(wi.Date, "/")
	day := tanggalInvoice[0]
	month, _ := strconv.Atoi(string(tanggalInvoice[1]))
	year := tanggalInvoice[2]
	monthMap := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	date := fmt.Sprintf("%s-%s-%s", day, monthMap[month], year)

	if len(wi.Products) == 0 {
		return InvoiceData{}, fmt.Errorf("invoice products is empty")
	}
	totalD := 0.0
	for _, v := range wi.Products {
		totalD = (float64(v.Unit) * v.Price) - v.Discount
	}

	dName := "warehouse"

	return InvoiceData{
		Date:         date,
		TotalInvoice: totalD,
		Departemen:   DepartmentName(dName),
	}, nil // TODO: replace this
}
