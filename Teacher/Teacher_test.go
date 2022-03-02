package Teacher

import (
	"testing"
	"week3/school/Course"
)

func TestTeachers_GetAttendancePercentage(t *testing.T) {
	test := []struct {
		input0 Teachers

		output int32
	}{
		{input0: Teachers{TeacherName: "Gbenga", TeacherAge: 35, YearsOfService: 4, Course: Course.Course{}, StudentScores: nil, StudentAttendance: nil}, output: 0},
	}
	for _, v := range test {
		result := v.input0.GetAttendancePercentage("Chuks")
		if result != v.output {
			t.Errorf("expected %v, got %v", v.output, result)
		}
	}
}
