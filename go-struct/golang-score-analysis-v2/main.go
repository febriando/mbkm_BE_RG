package main

import (
	"fmt"
)

func Add(a, b int) int {
	return 0
}

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {

	if s.Grades == nil || len(s.Grades) == 0 {
		return 0, 0, 0
	}

	var sum int
	min := s.Grades[0]
	max := s.Grades[0]

	for i := 0; i < len(s.Grades); i++ {
		if len(s.Grades) == 1 {
			max = s.Grades[i]
			min = s.Grades[i]
		}
		if max > s.Grades[i] {
			max = s.Grades[i]
		} else if min < s.Grades[i] {
			min = s.Grades[i]
		}
		sum += s.Grades[i]
	}

	avg := float64(sum) / float64(len(s.Grades))

	return avg, max, min // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60, 60, 100, 100, 100, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57},
	})

	fmt.Println(avg, min, max)
}
