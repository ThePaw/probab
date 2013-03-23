// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:betabinexch
func TestBetaBinExch(t *testing.T) {
	fmt.Println("Testing BetaBinExch()")
	theta1 := -1.0
	theta2 := 0.0
	n := []float64{20, 20, 20, 20, 20}
	c := []float64{1, 4, 3, 6, 10}

	x := BetaBinExch(theta1, theta2, c, n)
	y := -58.4339

	if abs(x/y-1) > 1e4 {
		fmt.Println("means failed: x, y ", x, y)
		t.Error()
	}
}

// Test against R:LearningBayes:betabinexch0
func TestBetaBinExch0(t *testing.T) {
	fmt.Println("Testing BetaBinExch0()")
	theta1 := 0.1
	theta2 := 10.0
	n := []float64{20, 20, 20, 20, 20}
	c := []float64{1, 4, 3, 6, 10}

	x := BetaBinExch(theta1, theta2, c, n)
	y := -58.82094

	if abs(x/y-1) > 1e4 {
		fmt.Println("means failed: x, y ", x, y)
		t.Error()
	}
}
