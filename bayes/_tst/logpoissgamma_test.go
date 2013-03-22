package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes
func TestLogPoissGamma(t *testing.T) {
	fmt.Println("TestLogPoissGamma #1")
	theta := []float64{-1, 0, 1, 2}
	d := []float64{2, 4, 3, 6, 1, 0, 4, 3, 10, 2}
	sh := 1.0
	rt := 1.0
	y := []float64{-49.289786, -20.243112, -3.144212, -18.522729}

	x := LogPoissGamma(theta, d, sh, rt)

	for i := range x {
		if abs(x[i]-y[i]) > 1e4 {
			fmt.Println("failed: i, x[i] , y[i] ", i, x[i], y[i])
			t.Error()
		}
	}

	fmt.Println("TestLogPoissGamma #2")
	sh = 1.4
	rt = 0.6
	y = []float64{-50.138177, -20.438655, -2.252443, -15.362650}

	x = LogPoissGamma(theta, d, sh, rt)

	for i := range x {
		if abs(x[i]-y[i]) > 1e4 {
			fmt.Println("failed: i, x[i] , y[i] ", i, x[i], y[i])
			t.Error()
		}
	}
}
