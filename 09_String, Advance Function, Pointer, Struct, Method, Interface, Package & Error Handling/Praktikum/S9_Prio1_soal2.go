package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func (s *Student) averageScore() float64 {
	totalScore := 0
	for _, score := range s.score {
		totalScore += score
	}
	return float64(totalScore) / float64(len(s.score))
}

func (s *Student) minScore() int {
	minScore := s.score[0]
	for _, score := range s.score {
		if score < minScore {
			minScore = score
		}
	}
	return minScore
}

func (s *Student) maxScore() int {
	maxScore := s.score[0]
	for _, score := range s.score {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	// membuat 5 siswa
	students := []Student{}
	for i := 1; i <= 5; i++ {
		var name string
		var score int
		fmt.Printf("Input %d Student's name: ", i)
		fmt.Scanln(&name)
		fmt.Printf("Input %d Student's Score: ", i)
		fmt.Scanln(&score)
		students = append(students, Student{name, []int{score}})
	}

	// mencari skor rata-rata
	totalScore := 0.0
	for _, student := range students {
		totalScore += student.averageScore()
	}
	averageScore := totalScore / float64(len(students))

	// mencari siswa dengan skor minimum
	minScore := students[0].minScore()
	var studentWithMinScore Student
	for _, student := range students {
		if student.minScore() == minScore {
			studentWithMinScore = student
			break
		}
	}

	// mencari siswa dengan skor maksimum
	maxScore := students[0].maxScore()
	var studentWithMaxScore Student
	for _, student := range students {
		if student.maxScore() == maxScore {
			studentWithMaxScore = student
			break
		}

	}

	// menampilkan hasil
	fmt.Printf("Skor rata-rata: %.2f\n", averageScore)
	fmt.Printf("Siswa dengan skor minimum: %s (skor: %d)\n", studentWithMinScore.name, minScore)
	fmt.Printf("Siswa dengan skor maksimum: %s (skor: %d)\n", studentWithMaxScore.name, maxScore)
}
