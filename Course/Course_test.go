package Course

import (
	"reflect"
	"testing"
)

func TestCourse_PrintCourseList(t *testing.T) {
	test := []struct {
		input1 Course

		output []string
	}{
		{Course{CourseName: "Geophysics", CourseCreditLoad: 30, CourseList: []string{"Physics", "Chemistry", "Math"}}, []string{"Physics", "Chemistry", "Math"}},
	}
	for _, v := range test {
		result := v.input1.PrintCourseList()
		if reflect.DeepEqual(result, v.output) != true {
			t.Errorf("expected %v, got %v", v.output, result)
		}
	}
}
