package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	// bonusJ := (1 * j.BaseSalary) * (j.WorkingMonth / 12)
	// return float64(bonusJ)
	return (1 * float64(j.BaseSalary)) * (float64(j.WorkingMonth) / 12)
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	// bonusS := (2 * s.BaseSalary * s.WorkingMonth / 12) + (int(s.PerformanceRate) * s.BaseSalary)
	// return float64(bonusS)
	return (2 * float64(s.BaseSalary) * float64(s.WorkingMonth) / 12) + (s.PerformanceRate * float64(s.BaseSalary))
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	// bonusM := (2 * m.BaseSalary * m.WorkingMonth / 12) + (int(m.PerformanceRate) * m.BaseSalary)
	// return float64(bonusM)
	return (2 * float64(m.BaseSalary) * float64(m.WorkingMonth) / 12) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalGajiEmployees := 0.0
	for _, v := range employees {
		totalGajiEmployees += v.GetBonus()
	}
	return totalGajiEmployees // TODO: replace this
}

func main() {
	kj := Junior{
		Name:         "Rian",
		BaseSalary:   100,
		WorkingMonth: 12,
	}
	ks := Senior{
		Name:            "Ping",
		BaseSalary:      100,
		WorkingMonth:    12,
		PerformanceRate: 0.5,
	}
	km := Manager{
		Name:             "Prop",
		BaseSalary:       100,
		WorkingMonth:     12,
		PerformanceRate:  0.5,
		BonusManagerRate: 0.5,
	}
	fmt.Println(EmployeeBonus(kj))
	fmt.Println(EmployeeBonus(ks))
	fmt.Println(EmployeeBonus(km))

	arrEmployees := []Employee{kj, ks, km}
	fmt.Println(TotalEmployeeBonus(arrEmployees))
}
