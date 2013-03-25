package bayes

import (
	"fmt"
	"testing"
)

func TestNormMeanTestTwoSided(t *testing.T) {
	fmt.Println("NormMeanTestTwoSided:  ")
	m0 := 170.0
	prob := 0.5
	smpMean := 176.0
	smpSize := 10
	popSd := 3.0
	tau := []float64{0.5, 1, 2, 4, 8}
	bf, post := NormMeanTestTwoSided(m0, prob, tau, smpMean, smpSize, popSd)
	fmt.Println("bf ", bf, post)
	fmt.Println("should be 1.462146e-02 3.897038e-05 1.894326e-07 2.591162e-08 2.309739e-08")
	fmt.Println("post ", post)
	fmt.Println("should be 1.441076e-02 3.896887e-05 1.894325e-07 2.591162e-08 2.309739e-08")
}
