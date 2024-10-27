package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) CalculateAverage() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	var sum float64
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

func main() {
	var student Student
	var numGrades int

	fmt.Print("Enter the student's name: ")
	fmt.Scanln(&student.Name)

	fmt.Print("Enter the number of grades: ")
	fmt.Scanln(&numGrades)

	student.Grades = make([]float64, numGrades)

	for i := 0; i < numGrades; i++ {
		fmt.Printf("Enter grade %d: ", i+1)
		fmt.Scanln(&student.Grades[i])
	}

	average := student.CalculateAverage()
	fmt.Printf("Average grade for %s: %.2f\n", student.Name, average)
}
