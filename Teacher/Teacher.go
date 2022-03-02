package Teacher

import "week3/school/Course"

type Teachers struct {
	TeacherName       string
	TeacherAge        int32
	YearsOfService    int32
	Course            Course.Course
	StudentScores     map[string]float64
	StudentAttendance map[string]int32
}

func (t *Teachers) GetDistinctionStudents(cutoffMark float64) map[string]float64 {
	distinctionStudent := make(map[string]float64)
	for k, v := range t.StudentScores {
		if v >= cutoffMark {
			distinctionStudent[k] = v
		}
	}
	return distinctionStudent
}
func (t *Teachers) GetFailingStudents(cutoffMark float64) map[string]float64 {
	failingStudents := make(map[string]float64)
	for k, v := range t.StudentScores {
		if v <= cutoffMark {
			failingStudents[k] = v
		}
	}
	return failingStudents
}

func (t *Teachers) GetAttendancePercentage(s string) int32 {
	if val, ok := t.StudentAttendance[s]; ok {
		return val
	}
	return 0
}
