package main

import (
	"os"
	"testing"
)

func TestSquareDetectionTrue(t *testing.T) {
	file, _ := os.Open("test_true.in")
	answers := Run(file)
	if len(answers) != 9 {
		t.Error("Incorrect number of answers")
	}
	for i := range answers {
		if answers[i] != true {
			t.Errorf("Unmatched case #%d (expected %s but got %s)", i+1, true, answers[i])
		}
	}
}

func TestSquareDetectionFalse(t *testing.T) {
	file, _ := os.Open("test_false.in")
	answers := Run(file)
	if len(answers) != 13 {
		t.Error("Incorrect number of answers")
	}
	for i := range answers {
		if answers[i] != false {
			t.Errorf("Unmatched case #%d (expected %s but got %s)", i+1, false, answers[i])
		}
	}
}
