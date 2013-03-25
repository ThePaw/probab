// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:logctablepost
func TestLogCTablePost(t *testing.T) {
	fmt.Println("Testing LogCTablePost()")
	s1, f1, s2, f2, theta1, theta2 := 6.0, 2.0, 3.0, 10.0, 2.0, 4.0
	x := LogCTablePost(s1, f1, s2, f2, theta1, theta2)
	y := -20.4611

	if abs(x/y-1) > 1e4 {
		fmt.Println("means failed: x, y ", x, y)
		t.Error()
	}
}
