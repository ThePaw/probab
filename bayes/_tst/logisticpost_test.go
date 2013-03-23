// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:logisticpost
func TestLogisticPost(t *testing.T) {
	fmt.Println("Testing LogisticPost()")
	beta0 := 2.0
	beta1 := 10.0

	d := []float64{-0.86, -0.3, -0.05, 0.73}
	n := []float64{5, 5, 5, 5}
	z := []float64{0, 1, 3, 5}
	x := LogisticPost(d, n, z, beta0, beta1)
	y := -6.580629

	if abs(x/y-1) > 1e4 {
		fmt.Println("means failed: x, y ", x, y)
		t.Error()
	}
}
