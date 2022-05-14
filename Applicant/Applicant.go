package Applicants

type Applicants struct {
	ApplicantName  string
	ApplicantAge   int32
	JambScore      float64
	CourseOfChoice string
	StateOfOrigin  string
}

func (app *Applicants) GetScore() float64 {
	return app.JambScore
}
func (app *Applicants) Admitted(score float64, age int32) string {
	if app.JambScore >= score && app.ApplicantAge >= age {
		return app.ApplicantName + " has gotten admission into the school"
	}
	return app.ApplicantName + " wasn't admitted"
}
