package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes:ctable(y,a)
func TestCTab(t *testing.T) {
	a := [][]float64{
		{1, 1, 1},
		{1, 1, 1},
	}

	y := [][]float64{
		{10, 6, 6},
		{4, 3, 10},
	}

	fmt.Println("Testing CTab #1")

	c1 := CTab(y, a)
	c2 := 2.288513
	if !check(c1, c2) {
		t.Error()
		fmt.Println("CTab: ", CTab(y, a), "should be ", c2)
	}

	fmt.Println("Testing CTab #2")
	a = [][]float64{
		{1, 1},
		{1, 1},
	}

	y = [][]float64{
		{10, 4},
		{4, 10},
	}
	c1 = CTab(y, a)
	c2 = 7.490698
	if !check(c1, c2) {
		t.Error()
		fmt.Println("CTab: ", CTab(y, a), "should be ", c2)
	}
}
