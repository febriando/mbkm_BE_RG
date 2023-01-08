package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total_minjam"`
}

func FindEmployee(id string, employees []Employee) (Employee, bool) {
	for _, employee := range employees {
		if employee.ID == id {
			return employee, true
		}
	}

	return Employee{}, false
}

func GetEndBalanceAndBorrowers(data LoanData) (int, []Borrower) {
	if len(data.Data) != 0 {
		// recap the data in map
		var tempMap = map[string]int{}

		for _, v := range data.Data {
			for _, id := range v.EmployeeIDs {
				if data.StartBalance >= 50000 {
					tempMap[id] += 50000
					data.StartBalance -= 50000
				}
			}
		}

		// sort temp
		var keys []string

		for k := range tempMap {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		var borrowers []Borrower

		// set data to struct
		for _, k := range keys {
			if employee, ok := FindEmployee(k, data.Employees); ok {
				borrowers = append(borrowers, Borrower{
					ID:    employee.ID,
					Name:  employee.Name,
					Total: tempMap[k],
				})
			}
		}

		// set the start balance and end balance
		return data.StartBalance, borrowers
	} else {
		return 0, []Borrower{}
	}
}

func LoanReport(data LoanData) LoanRecord {
	// return LoanRecord{} // TODO: replace this

	endBalance, borrowers := GetEndBalanceAndBorrowers(data)
	lData := LoanData{}
	lRecord := LoanRecord{}
	if lRecord.StartBalance < lData.StartBalance {
		endBalance = 0
	}
	lRecord.EndBalance = endBalance
	lRecord.StartBalance = data.StartBalance
	lRecord.Borrowers = append(lRecord.Borrowers, borrowers...)

	// fmt.Println(lRecord)

	return lRecord
}

func RecordJSON(record LoanRecord, path string) error {
	// return nil // TODO: replace this
	json, err := json.Marshal(record)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, json, 0644)
	if err != nil {
		return err
	}

	return err
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}