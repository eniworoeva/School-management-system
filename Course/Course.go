package Course

import "fmt"

type Course struct {
	CourseName       string
	CourseCreditLoad int32
	CourseList       []string
}

func (c *Course) PrintCourseList() []string {
	for _, val := range c.CourseList {
		fmt.Printf("%s\n", val)
	}
	return c.CourseList
}
