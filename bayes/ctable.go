// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Bayes factor against independence in a 
// 2-way contingency table using uniform priors.
// Ref.: Albert (2009)

import (
	fn "code.google.com/p/go-fn/fn"
	"math"
)

func lDirich(a []float64) float64 {
	sumLnΓ := 0.0
	sum := 0.0
	for _, x := range a {
		sumLnΓ += fn.LnΓ(x)
		sum += x
	}
	return sumLnΓ - fn.LnΓ(sum)
}

func rowSums(a [][]float64) []float64 {
	v := make([]float64, len(a))
	for i, row := range a {
		for _, x := range row {
			v[i] += x
		}
	}
	return v
}

func colSums(a [][]float64) []float64 {
	v := make([]float64, len(a[0]))
	for _, row := range a {
		for j, x := range row {
			v[j] += x
		}
	}
	return v
}

// FactCTableUnif returns the Bayes factor for testing independence in a contingency table.
func FactCTableUnif(y, a [][]float64) float64 {
	// Bayes factor against independence in a 
	// 2-way contingency table using uniform priors, 
	// where Y is a matrix containing the 2-way table of counts, 
	// and A is a matrix of prior hyperparameters.

	nRows := len(y)
	nCols := len(y[0])
	fRows := float64(nRows)
	fCols := float64(nCols)
	ac := colSums(a)
	ar := rowSums(a)
	yc := colSums(y)
	yr := rowSums(y)

	t1 := make([]float64, nRows*nCols)
	k := 0
	for j := 0; j < nCols; j++ {
		for i := 0; i < nRows; i++ {
			t1[k] = y[i][j] + a[i][j]
			k++
		}
	}

	t2 := make([]float64, nRows)
	for i := 0; i < nRows; i++ {
		t2[i] = ar[i] - fCols + 1
	}

	t3 := make([]float64, nCols)
	for i := 0; i < nCols; i++ {
		t3[i] = ac[i] - fRows + 1
	}

	t4 := make([]float64, nRows*nCols)
	k = 0
	for j := 0; j < nCols; j++ {
		for i := 0; i < nRows; i++ {
			t4[k] = a[i][j]
			k++
		}
	}

	t5 := make([]float64, nRows)
	for i := 0; i < nRows; i++ {
		t5[i] = yr[i] + ar[i] - fCols + 1
	}

	t6 := make([]float64, nCols)
	for i := 0; i < nCols; i++ {
		t6[i] = yc[i] + ac[i] - fRows + 1
	}

	lbf := lDirich(t1) + lDirich(t2) + lDirich(t3) - lDirich(t4) - lDirich(t5) - lDirich(t6)
	return math.Exp(lbf)
}
