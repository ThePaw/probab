// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Some utility functions.

// meanSd returns mean and standard deviation of a vector.
func meanSd(data []float64) (mean, sd float64) {
	n := 0.0
	mean = 0.0
	m2 := 0.0
	for _, x := range data {
		n++
		delta := x - mean
		mean += delta / n
		if n > 1 {
			m2 += delta * (x - mean)
		}
	}
	sd = sqrt(m2 / (n - 1))
	return
}

// cols2vec unloads complete matrix y columnwise to vector v.
func cols2vec(y [][]float64) []float64 {
	v := make([]float64, len(y)*len(y[0]))
	col := 0
	for j, _ := range y[0] {
		for i, _ := range y {
			v[col] = y[i][j]
			col++
		}
	}
	return v
}

// linInt is a linear interpolation function.
func linInt(x, y []float64, xVal float64) float64 {
	// x vector must be nondecreasing

	if xVal < x[0] || xVal > x[len(x)-1] {
		return nan
	}

	// find out which segment we are in
	n := 0
	for i, _ := range x {
		if xVal > x[i] {
			n = i
		} else {
			break
		}
	}
	if n > len(y) {
		return nan
	}

	// interpolate
	x0, y0 := x[n], y[n]
	x1, y1 := x[n+1], y[n+1]
	a := x1 - x0
	b := y1 - y0
	yVal := b*(xVal-x0)/a + y0
	return yVal
}
