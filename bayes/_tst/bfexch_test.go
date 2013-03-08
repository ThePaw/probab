package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:bfexch()
func TestBFExch(t *testing.T) {
	y := []float64{1, 3, 2, 4, 6, 4, 3}
	n := []float64{10, 10, 10, 10, 10, 10, 10}
	theta := 1.0
	k := 20.0

	fmt.Println("Testing BFExch #1")

	c1 := BFExch(theta, y, n, k)
	c2 := -16.48318
	if !check(c1, c2) {
		t.Error()
		fmt.Println("BFExch: ", c1, "should be ", c2)
	}
	fmt.Println("BFExch: ", c1, "should be ", c2)

}
