package bayes

import (
	"fmt"
	"testing"
)

func TestNormMeanTestOneSided(t *testing.T) {
	fmt.Println("NormMeanTestOneSided:  ")
	bf, priOdds, postOdds, postH := NormMeanTestOneSided(175, 170, 1000, 176, 10, 3)
	fmt.Println("bf, priOdds, postOdds, postH", bf, priOdds, postOdds, postH)
	fmt.Println("should be 0.1694947 1.008011 0.1708525 0.1459215")
}

// Ref.: Albert (2009)
