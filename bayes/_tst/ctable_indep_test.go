package bayes

import (
	"code.google.com/p/probab/dst"
	"fmt"
	"testing"
)

// Test against R:LearningBayes:bfindep()
func TestFactCTableIndep(t *testing.T) {
	y := [][]float64{
		{1, 10},
		{5, 20},
	}
	k := 20.0
	m := 1000000
	fmt.Println("Testing ldirichlet")

	c5 := ldirichlet(y)
	fmt.Println("ldirichlet: ", c5)
	fmt.Println("should be: -2.302585 -12.266791")

	fmt.Println("Testing meanSd")

	nrm := make([]float64, m)
	for i, _ := range nrm {
		nrm[i] = dst.NormalNext(5, 0.6)
	}
	mu, sd := meanSd(nrm)
	fmt.Println("mu, sd :", mu, sd)

	fmt.Println("cols2vec")
	fmt.Println(cols2vec(y))
	fmt.Println("should be 1 5 10 20")

	fmt.Println("Testing FactCTableIndep #1")
	c1, c2 := FactCTableIndep(y, k, m)
	c3, c4 := 0.6768298, 0.002508772
	if !check(c1, c3) || !check(c2, c4) {
		t.Error()
		fmt.Println("FactCTableIndep: bf = ", c1, "should be ", c3, "+/- 0.015")
		fmt.Println("                nse = ", c2, "should be ", c4, "+/- 0.003")
	}

}
