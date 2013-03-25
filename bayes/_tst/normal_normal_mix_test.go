// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:normal.normal.mix
func TestNormalNormalMix(t *testing.T) {
	fmt.Println("Testing NormalNormalMix()")
	probs := []float64{.5, .5}
	priorMean := []float64{0.0, 2.0}
	priorVar := []float64{1.0, 0.5}
	y := 1.0
	sigma2 := .5

	postProbs, postMean, postVar := NormalNormalMix(probs, priorMean, priorVar, y, sigma2)

	fmt.Println("postProbs ", postProbs)
	fmt.Println("should be: [0.4909845   0.5090155]")
	fmt.Println("postMean ", postMean)
	fmt.Println("should be: [0.6666667 1.5000000]")
	fmt.Println("postVar ", postVar)
	fmt.Println("should be: [0.3333333 0.2500000]")
}
