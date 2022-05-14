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

func (t *Teachers) GetAttendancePercentage(s string) int32 {
	if val, ok := t.StudentAttendance[s]; ok {
		return val
	}
	return 0
}
