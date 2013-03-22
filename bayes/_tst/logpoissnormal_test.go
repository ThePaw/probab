package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes
func TestLogPoissNormal(t *testing.T) {
	theta := []float64{-1, 0, 1, 2}
	d := []float64{2, 4, 3, 6, 1, 0, 4, 3, 10, 2}
	mean := 0.0
	sd := 1.0

	y := []float64{-49.340845, -20.162051, -2.844869, -16.052612}

	x := LogPoissNormal(theta, d, mean, sd)

	for i := range x {
		if abs(x[i]-y[i]) > 1e4 {
			fmt.Println("failed: i, x[i] , y[i] ", i, x[i], y[i])
			t.Error()
		}
	}
}
