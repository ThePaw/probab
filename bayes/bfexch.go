// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Logarithm of the integral of the Bayes factor for testing homogeneity of a set of proportions.
// Ref.: Albert (2009)

import (
	fn "code.google.com/p/go-fn/fn"
	"math"
)

var lnB func(float64, float64) float64 = fn.LnB
var exp func(float64) float64 = math.Exp
var sqrt func(float64) float64 = math.Sqrt
var log func(float64) float64 = math.Log

func logF(k, eta float64, y, n []float64) []float64 {
	v := make([]float64, len(y))
	for i, _ := range v {
		v[i] = lnB(k*eta+y[i], k*(1-eta)+n[i]-y[i]) - lnB(k*eta, k*(1-eta))
	}
	return v
}

// BFExch returns the logarithm of the integral of the Bayes factor for testing homogeneity 
// of a set of proportions.
func BFExch(theta float64, y, n []float64, k float64) float64 {
	eta := exp(theta) / (1 + exp(theta))
	sum := 0.0
	sumY := 0.0
	sumNY := 0.0
	w := logF(k, eta, y, n)
	for i, _ := range y {
		sum += w[i]
		sumY += y[i]
		sumNY += n[i] - y[i]
	}
	return sum + log(eta*(1-eta)) - lnB(sumY+1, sumNY+1)
}
