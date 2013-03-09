package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes
func TestBetaprior1(t *testing.T) {
	fmt.Println("betaprior1 ", betaprior1(5, 0.5, 0.45))
	fmt.Println("should be 0.5258179")
}

func TestBetaSelect(t *testing.T) {
	x, y := BetaFromQtls(0.5, 0.25, 0.9, 0.45)
	fmt.Println("BetaFromQtls ", x, y)
	fmt.Println("should be 2.67 7.37")
}
