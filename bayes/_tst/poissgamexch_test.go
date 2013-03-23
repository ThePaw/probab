// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:poissgamexch
func TestPoissGamExch(t *testing.T) {
	fmt.Println("Testing PoissGamExch()")
	theta1 := -4.0
	theta2 := 0.0
	z0 := 0.5

	e := []float64{532, 584, 672, 722, 904}
	c := []float64{0, 0, 2, 1, 1}

	x := PoissGamExch(theta1, theta2, z0, c, e)
	y := -42.03939

	if abs(x/y-1) > 1e4 {
		fmt.Println("means failed: x, y ", x, y)
		t.Error()
	}
}
