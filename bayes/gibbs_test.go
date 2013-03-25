// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"fmt"
	"testing"
)
func logpost(theta []float64) float64 {
	s1, f1, s2, f2 := 6.0, 2.0, 3.0, 10.0
	return LogCTablePost(s1, f1, s2, f2, theta[0], theta[1])
}

// Test against R:LearningBayes:gibbs
func TestGibbs(t *testing.T) {
	m := 10000000
	start := []float64{1.0, 1.0}
	scale := []float64{2.0, 2.0}
	fmt.Println("Testing Gibbs")
	_, arate := Gibbs(logpost, start,  m , scale )
	fmt.Println("arate : ",  arate)
}
