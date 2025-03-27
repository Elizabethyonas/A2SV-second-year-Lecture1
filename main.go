package main

import "fmt"

func calculateAverage(grades []float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func main() {
	var studentName string
	var numberOfSubjects int

	fmt.Print("Enter student name: ")
	fmt.Scanln(&studentName)

	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&numberOfSubjects)

	if numberOfSubjects <= 0 {
		fmt.Println("Invalid number of subjects")
		return
	}

	subjects := make(map[string]float64)
	grades := make([]float64, 0, numberOfSubjects)

	for i := 0; i < numberOfSubjects; i++ {
		var subjectName string
		var grade float64

		fmt.Print("Enter subject name: ")
		fmt.Scanln(&subjectName)

		fmt.Print("Enter grade: ")
		fmt.Scanln(&grade)

		subjects[subjectName] = grade
		grades = append(grades, grade)
	}

	average := calculateAverage(grades)

	fmt.Printf("Student: %s\n", studentName)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Println("Subjects:")
	for subject, grade := range subjects {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
}
