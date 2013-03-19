package bayes

import (
	"fmt"
	"math/rand"
	"testing"
)

// Test against R:LearningBayes
func TestLnHowardPrior(t *testing.T) {
	z := LnHowardPrior(0.1, 0.5, 1, 1, 1, 1, 2)
	fmt.Println("LnHowardPrior :", z)
	fmt.Println("should be -0.6034745")
}
func TestHowardPosteriorProb(t *testing.T) {
	rand.Seed(56589856)
	sigma := 2.0
	z := HowardPosteriorProb(3, 18, 7, 12, 1, 1, 1, 1, sigma)
	fmt.Println("HowardPosteriorProb :", z)
	fmt.Println("should be 0.0161")
	sigma = 1.0
	z = HowardPosteriorProb(3, 18, 7, 12, 1, 1, 1, 1, sigma)
	fmt.Println("HowardPosteriorProb :", z)
	fmt.Println("should be 0.0343")
	sigma = 0.5
	z = HowardPosteriorProb(3, 18, 7, 12, 1, 1, 1, 1, sigma)
	fmt.Println("HowardPosteriorProb :", z)
	fmt.Println("should be 0.1037")
	sigma = 0.25
	z = HowardPosteriorProb(3, 18, 7, 12, 1, 1, 1, 1, sigma)
	fmt.Println("HowardPosteriorProb :", z)
	fmt.Println("should be 0.2374")
}
