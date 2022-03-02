package Course

import "fmt"

type Course struct {
	CourseName       string
	CourseCreditLoad int32
	CourseList       []string
}

func (c *Course) PrintCourseList() []string {
	for ind, val := range c.CourseList {
		fmt.Printf("%d : %s\n", ind+1, val)
	}
	return c.CourseList
}

//func (c *Course) CheckCourseInList(myCourse string) bool {
//	for i := 0; i < len(c.CourseList); i++ {
//		if c.CourseList[i] == myCourse{
//			return true
//		}
//	}
//	return false
//}
//func (c *Course) AddNewCourse(newCourse string) {
//	if isPresent := c.CheckCourseInList(newCourse); !isPresent {
//		c.CourseList = append(c.CourseList, newCourse)
//	}
//}
