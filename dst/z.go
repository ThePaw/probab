// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Standard Normal  (or Gaussian, or Gauss-Laplace) distribution. 
// A continuous probability distribution, defined on the entire real line, that has a bell-shaped probability density function, known as the Gaussian function. 
//
// Parameters: none
//
// Support: 
// x ∈ R

import (
	"math"
//	"math/rand"
)


// Z_PDF returns the PDF of the Standard Normal distribution. 
func Z_PDF() func(float64) float64 {
	return Normal_PDF(0, 1)
}

// Z_PDF_At returns the value of PDF of Standard Normal distribution at x. 
func Z_PDF_At(x float64) float64 {
	pdf := Normal_PDF(0, 1)
	return pdf(x)
}

// Z_CDF returns the CDF of the Standard Normal distribution. 
func Z_CDF() func(float64) float64 {
	return Normal_CDF(0, 1)
}

// Z_CDF_At returns the value of CDF of the Standard Normal distribution, at x. 
func Z_CDF_At(x float64) float64 {
	cdf := Normal_CDF(0, 1)
	return cdf(x)
}

// Z_Qtl returns the inverse of the CDF (quantile) of the Standard Normal distribution. 
func Z_Qtl() func(p float64)  float64 {
	return func(p float64)  float64 {

	var r, x, pp, dp float64

	dp = p - 0.5
	switch {
	case p == 1.0:
		return math.MaxFloat64
	case p == 0.0:
		return -math.MaxFloat64
	}
	if math.Abs(dp) <= 0.425 {
		x = small(dp)
		return x
	}
	if p < 0.5 {
		pp = p
	} else {
		pp = 1.0 - p
	}
	r = math.Sqrt(-math.Log(pp))
	if r <= 5.0 {
		x = intermediate(r)
	} else {
		x = tail(r)
	}
	if p < 0.5 {
		return -x
	}
	return x
	}
}

// Z_Qtl_For returns the inverse of the CDF (quantile) of the Standard Normal distribution, for given probability.
func Z_Qtl_For(p float64) float64 {
	qtl := Z_Qtl()
	return qtl(p)
}

