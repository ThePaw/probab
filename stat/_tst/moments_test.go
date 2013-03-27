package stat

import (
	"code.google.com/p/probab/dst"
	"fmt"
	"testing"
)

// Test against R:moments
func TestSampleSkewness(t *testing.T) {
	fmt.Println("Testing Skewness")
	m := 10000000
	mu := 0.0
	sd := 1.0
	d := make([]float64, m)
	for i, _ := range d {
		d[i] = dst.NormalNext(mu, sd)
	}

	x := Skew(d)
	y := 0.0
	if abs(x-y) > 1e-3 {
		fmt.Println("failed: x, y ", x, y)
		t.Error()
	}
}

func TestSampleKurtosis(t *testing.T) {
	fmt.Println("Testing Kurtosis")
	m := 10000000
	mu := 0.0
	sd := 1.0
	d := make([]float64, m)
	for i, _ := range d {
		d[i] = dst.NormalNext(mu, sd)
	}

	x := Kurt(d)
	y := 3.0
	if abs(x-y) > 1e-3 {
		fmt.Println("failed: x, y ", x, y)
		t.Error()
	}
}

func TestSampleMoments(t *testing.T) {
	fmt.Println("Testing Moments")
	m := 10000000
	mu := 0.0
	sd := 1.0
	d := make([]float64, m)
	for i, _ := range d {
		d[i] = dst.NormalNext(mu, sd)
	}

	order := 4
	central := false
	absolute := false
	x := moment(d, order, central, absolute)
	y := 3.0
	if abs(x-y) > 1e-2 {
		fmt.Println("failed: x, y ", x, y)
		t.Error()
	}
}

func TestSampleGeary(t *testing.T) {
	fmt.Println("Testing Geary kurtosis")
	m := 1000000
	mu := 0.0
	sd := 1.0
	d := make([]float64, m)
	for i, _ := range d {
		d[i] = dst.NormalNext(mu, sd)
	}

	x := Geary(d)
	y := 0.7979113
	if abs(x-y) > 1e-2 {
		fmt.Println("failed: x, y ", x, y)
		t.Error()
	}
}
