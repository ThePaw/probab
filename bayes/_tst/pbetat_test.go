package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes
func TestProbBetaTest(t *testing.T) {
	p0 := .5
	prob := .5
	a := 10.0
	b := 10.0
	succ := 5
	fail := 15

	bf, post := ProbBetaTest(p0, prob, a, b, succ, fail)
	fmt.Println("ProbBetaTest bf, post:  ", bf, post)
	fmt.Println("should be 0.3893163, 0.2802215")
}
func TestCumSum(t *testing.T) {
	y := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("TestCumSum:  ", cumSum(y))
	fmt.Println("should be 1  3  6 10 15 21 28 36 45 55")
}
