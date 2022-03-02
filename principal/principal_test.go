package Principal

import (
	"fmt"
	"testing"
	Applicants "week3/school/Applicant"
	"week3/school/Student"
)

func TestPrincipal_AdmitStudent(t *testing.T) {
	n := recover()
	if n != nil {
		fmt.Println(n)
	}
	p := &Principal{"Francis", []string{"Denis", "carl", "vincent"}, map[string]string{}}
	table := []struct {
		jambite           Applicants.Applicants
		cutoffMark        float64
		specialCutOffMark float64
		output            bool
	}{
		{Applicants.Applicants{"Denis", 23, 350,
			"History", "Imo"}, 350, 300, true},
		{Applicants.Applicants{"John", 25, 310, "Medicine", "Ogun"}, 350, 380, false},
	}
	for i := 0; i < len(table); i++ {
		iterable := table[i]
		p.AdmitStudent(iterable.jambite, iterable.cutoffMark, iterable.specialCutOffMark)
		if _, ok := p.StudentList[iterable.jambite.ApplicantName]; ok {
			if ok != iterable.output {
				t.Errorf("AdmitStudent index %d : Expected %v  Got %v", i, iterable.output, ok)
			}
		}
	}
}
func TestPrincipal_ExpelStudent(t *testing.T) {
	p := Principal{"Francis", []string{"James", "fishing", "Tim", "Philosophy"}, map[string]string{}}
	table := []struct {
		input Student.Student

		expected bool
	}{
		{Student.Student{"James", "History", 250, []string{""}, true}, false},
		{Student.Student{"Tim", "Philosophy", 350, []string{""}, false}, true},
	}
	for _, val := range table {
		p.ExpelStudent(val.input)
		if _, ok := p.StudentList[val.input.StudentName]; ok {
			if val.input.PaidFees != ok {
				t.Errorf("Expected %v : Got %v", val.expected, ok)
			}
		}
	}
}
func TestPrincipal_GetStudentList(t *testing.T) {
	var testGetstudent = []struct {
		input  Principal
		input1 Student.Student
		output bool
	}{{Principal{
		"Francis",
		[]string{"glory", "history", "ijeoma", "Civil", "Adaobi", "Biology"}, map[string]string{},
	}, Student.Student{
		"paul", "englis", 330, []string{"history", "Biology"}, true,
	}, false},
	}

	for _, val := range testGetstudent {
		m := val.input.ExpelStudent(val.input1)
		if val.output != m {
			t.Errorf("Expected %v : Got %v", m, val.output)
		}
	}
}
