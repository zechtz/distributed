package grades

import (
	"fmt"
	"sync"
)

// Student defines a student struct
type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

// GradeType defines a grade type
type GradeType string

// Students is a slice of students
type Students []Student

var (
	students      Students
	studentsMutex sync.Mutex
)

const (
	// GradeTest constant type
	GradeTest = GradeType("Test")
	// GradeHowmework constant type
	GradeHowmework = GradeType("Homework")
	// GradeQuiz onstant type
	GradeQuiz = GradeType("Quiz")
)

// Grade define the structure of the grades object
type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

// Average returns the average grade score of a student
func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

// GetByID returns a given student based on id
func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}

	return nil, fmt.Errorf("Student with ID %v not found", id)
}
