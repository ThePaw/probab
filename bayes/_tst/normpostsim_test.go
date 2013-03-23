package bayes

import (
	"fmt"
	"testing"
)

// Test against R:LearningBayes
func TestNormPostSimNoPrior(t *testing.T) {
	fmt.Println("Testing NormPostSimNoPrior()")
	m := 100000
	d := []float64{-67, -48, 6, 8, 14, 16, 23, 24, 28, 29, 41, 49, 67, 60, 75}
	means, vars := NormPostSimNoPrior(d, m)

	x := mean(means)
	y := 21.67034

	if abs(x/y-1) > 1e4 {
		fmt.Println("means failed: x, y ", x, y)
		t.Error()
	}

	x = mean(vars)
	y = 1735.703

	if abs(x/y-1) > 1e4 {
		fmt.Println("vars failed: x, y ", x, y)
		t.Error()
	}
}

func TestNormPostNoPriorNext(t *testing.T) {
	fmt.Println("Testing NormPostNoPriorNext()")
	d := []float64{-67, -48, 6, 8, 14, 16, 23, 24, 28, 29, 41, 49, 67, 60, 75}
	m := 100000
	mn := 0.0
	vr := 0.0

	for i := 0; i < m; i++ {
		mni, vri := NormPostNoPriorNext(d)
		mn += mni
		vr += vri
	}
	mn /= float64(m)
	vr /= float64(m)

	y := 21.67034

	if abs(mn/y-1) > 1e4 {
		fmt.Println("means failed: mn, y ", mn, y)
		t.Error()
	}

	y = 1735.703

	if abs(vr/y-1) > 1e4 {
		fmt.Println("vars failed: vr, y ", vr, y)
		t.Error()
	}
}
