package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		fmt.Println("lulus")
		return "lulus"
	} else if score < 0 && score > 100 {
		fmt.Println("error, masukkan nilai score 0-100")
	} else if absent < 0 && absent > 10 {
		fmt.Println("error, masukkan nilai absent 0-10")
	} else {
		fmt.Println("tidak lulus")
		return "tidak lulus"
	}

	return "" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// var score, absent int
	// fmt.Print("Input Score: ")
	// fmt.Scanln(&score)
	// fmt.Print("Input Absent: ")
	// fmt.Scanln(&absent)

	// GraduateStudent(score, absent)

	GraduateStudent(70, 4)
	GraduateStudent(80, 1)
	GraduateStudent(70, 1)
	GraduateStudent(90, 0)
	GraduateStudent(95, 0)
	GraduateStudent(100, 0)
	GraduateStudent(69, 4)
	GraduateStudent(40, 1)
	GraduateStudent(69, 6)
	GraduateStudent(70, 10)
	GraduateStudent(100, 5)
	GraduateStudent(80, 6)
	GraduateStudent(0, 0)
	GraduateStudent(0, 6)
	GraduateStudent(0, 7)
	GraduateStudent(0, 8)
}
