package Student

import (
	"fmt"
)

type Student struct {
	StudentName string
	Department  string
	Scores      float64
	CourseName  []string
	PaidFees    bool
}

func (s Student) SackPrincipal() {
	fmt.Println("hello")
}
