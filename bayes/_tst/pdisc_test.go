package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes
func TestPropDisc(t *testing.T) {
	p := []float64{.2, .25, .3, .35}
	prior := []float64{.25, .25, .25, .25}
	succ := 5
	fail := 10

	out := PropDisc(p, prior, succ, fail)
	fmt.Println("PropDisc ", out)
	fmt.Println("should be 0.1502369 0.2404581 0.3001328 0.3091723")
}
