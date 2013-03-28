// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Highest density iterval (HDI) limits from quantile function, for Beta distribution.
// Ref: Kruschke 2012: Chapter 23.3.3, p. 629 and further.

import (
	"code.google.com/p/probab/dst"
)

// Interval width
func iw(α, β, credMass, lowTailPr float64) float64 {
	return dst.BetaQtlFor(α, β, credMass+lowTailPr) - dst.BetaQtlFor(α, β, lowTailPr)
}

// Interval width for fixed α, β, credMass
func iwFix(α, β, credMass float64) func(x float64) float64 {
	return func(x float64) float64 { return iw(α, β, credMass, x) }
}

// BetaHDI returns the Highhest Density Interval limits of the Beta Distribution.
func BetaHDI(α, β, credMass, tol float64) (lo, hi float64) {
	f := iwFix(α, β, credMass)
	// func fmin(f func(float64) float64, ax, bx, tol float64) float64 {

	min := fmin(f, 0, 1-credMass, tol)
	lo = dst.BetaQtlFor(α, β, min)
	hi = dst.BetaQtlFor(α, β, credMass+min)
	return lo, hi
}
