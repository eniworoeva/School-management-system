package Principal

import (
	"week3/school/Applicant"
	"week3/school/Student"
)

type Principal struct {
	PrincipalName string
	FinalList     []string
	StudentList   map[string]string //name
}

func (p *Principal) AdmitStudent(jambite Applicants.Applicants, cutOffMark float64, specialCutOffMark float64) {
	if jambite.JambScore >= cutOffMark {
		p.StudentList[jambite.ApplicantName] = jambite.CourseOfChoice
	} else if jambite.StateOfOrigin == "Lagos" && jambite.JambScore >= specialCutOffMark {
		p.StudentList[jambite.ApplicantName] = jambite.CourseOfChoice
	}
}
func (p *Principal) ExpelStudent(student Student.Student) bool {
	if student.PaidFees == false {
		return true
	}
	return false
}
