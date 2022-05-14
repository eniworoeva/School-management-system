package main

import (
	"fmt"
	Applicants "week3/school/Applicant"

	//"week3/school/Applicant"
	"week3/school/Course"
	"week3/school/Teacher"
)

func main() {

	c := Course.Course{CourseName: "Geophysics", CourseCreditLoad: 30, CourseList: []string{"Physics", "Chemistry", "Math"}}
	fmt.Println(c.PrintCourseList())

	d := Teacher.Teachers{TeacherName: "Gbenga", TeacherAge: 35, YearsOfService: 4, Course: Course.Course{}, StudentScores: nil, StudentAttendance: nil}
	fmt.Println(d.GetAttendancePercentage("Timi"))

	e := Applicants.Applicants{ApplicantName: "Segun", ApplicantAge: 20, JambScore: 239, CourseOfChoice: "Forestry", StateOfOrigin: "Lagos"}
	fmt.Println(e.GetScore())
	//fmt.Println(e.Admitted float64(256), int32(27))
}
