package Applicants

import "testing"

func TestApplicants_Admitted(t *testing.T) {
	admission := []struct {
		firstPointer Applicants
		input1       float64
		input2       int32
		output       string
	}{
		{Applicants{"Johnson", 14, 245.6, "Microbiology", "Benue"}, 250.7, 13, "Johnson wasn't admitted"},
		{Applicants{"Alex", 20, 272.6, "Forestry", "Enugu"}, 250.7, 13, "Alex has gotten admission into the school"},
	}
	for _, val := range admission {
		result := val.firstPointer.Admitted(val.input1, val.input2)
		if result != val.output {
			t.Errorf("expected %v, got %v", val.output, result)
		}
	}
}
