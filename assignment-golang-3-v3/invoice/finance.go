package invoice

import (
	"fmt"
	"strconv"
	"strings"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	if len(fi.Date) == 0 {
		return InvoiceData{}, fmt.Errorf("invoice date is empty")
	}
	tanggalInvoice := strings.Split(fi.Date, "/")
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

	if len(fi.Details) == 0 {
		return InvoiceData{}, fmt.Errorf("invoice details is empty")
	}
	totalD := 0
	for _, p := range fi.Details {
		totalD += p.Total
	}
	if totalD <= 0 {
		return InvoiceData{}, fmt.Errorf("total price is not valid")
	}

	if fi.Status != InvoiceStatus("paid") || fi.Status != InvoiceStatus("unpaid") || fi.Status == "" {
		return InvoiceData{}, fmt.Errorf("invoice status is not valid")
	}

	dName := "finance"

	return InvoiceData{
		Date:         date,
		TotalInvoice: float64(totalD),
		Departemen:   DepartmentName(dName),
	}, nil // TODO: replace this
}
