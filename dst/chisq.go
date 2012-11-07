// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Chi-Squared distribution. 
// Parameters: 
// n ∈ ℕ	(degrees of freedom)
// Support: 
// x ∈ [0, +∞]

import (
	"math"
	. "code.google.com/p/go-fn/fn"
)

// ChiSquare_PDF returns the PDF of the ChiSquare distribution. 
func ChiSquare_PDF(n int64) func(x float64) float64 {
	k := float64(n) / 2
	normalization := pow(0.5, k) / Γ(k)
	return func(x float64) float64 {
		return normalization * pow(x, k-1) * NextExp(-x/2)
	}
}

// ChiSquare_LnPDF returns the natural logarithm of the PDF of the ChiSquare distribution. 
func ChiSquare_LnPDF(n int64) func(x float64) float64 {
	k := float64(n) / 2
	normalization := log(0.5)*k - LnΓ(k)
	return func(x float64) float64 {
		return normalization + log(x)*(k-1) - x/2
	}
}
// ChiSquare_PDF_At returns the value of PDF of ChiSquare distribution at x. 
// Uniform_PDF_At returns the value of PDF of Uniform distribution(μ, σ) at x. 
func ChiSquare_PDF_At(n int64, x float64) float64 {
	pdf := ChiSquare_PDF(n)
	return pdf(x)
}

// ChiSquare_CDF returns the CDF of the ChiSquare distribution. 
func ChiSquare_CDF(n int64) func(x float64) float64 {
	return func(x float64) float64 {
		return Γr(float64(n)/2, x/2)
	}
}

// ChiSquare_CDF_At returns the value of CDF of the ChiSquare distribution, at x. 
func ChiSquare_CDF_At(n int64, x float64) float64 {
	cdf := ChiSquare_CDF(n)
	return cdf(x)
}

// ChiSquare_Qtl returns the inverse of the CDF (quantile) of the ChiSquare distribution. 
func ChiSquare_Qtl(n int64) func(p float64) float64 {
	return func(p float64) float64 {
		//return Gamma_Qtl_At(n/2, 2, p)  to be implemented
		return Gamma_Qtl_For(float64(n)/2, 2, p)
	}
}

// NextChiSquare returns random number drawn from the ChiSquare distribution. 
func NextChiSquare(n int64) (x float64) {
//ChiSquare(n) => sum of n N(0,1)^2
	for i := iZero; i < n; i++ {
		n := NextNormal(0, 1)
		x += n * n
	}
	return
}

// ChiSquare returns the random number generator with  ChiSquare distribution. 
func ChiSquare(n int64) func() float64 {
	return func() float64 {
		return NextChiSquare(n)
	}
}

// ChiSquareMean returns the mean of the ChiSquare distribution. 
func ChiSquareMean(n int64) float64 {
	return float64(n)
}

// ChiSquareMedian returns the approximate median of the ChiSquare distribution. 
func ChiSquareMedian(n int64) float64 {
	c:=1-(2.0/(9.0*n))
	c = c*c*c
	return float64(n *c)
}

// ChiSquareMode returns the mode of the ChiSquare distribution. 
func ChiSquareMode(n int64) float64 {
	return math.Max(float64(n-2), 0)
}

// ChiSquareVar returns the variance of the ChiSquare distribution. 
func ChiSquareVar(n int64) float64 {
	return float64(2*n)
}

// ChiSquareStd returns the standard deviation of the ChiSquare distribution. 
func ChiSquareStd(n int64) float64 {
	return math.Pow(float64(2*n), 0.5)
}

// ChiSquareSkew returns the skewness of the ChiSquare distribution. 
func ChiSquareSkew(n int64) float64 {
	return math.Pow(float64(8/n), 0.5)
}

// ChiSquareExKurt returns the excess kurtosis of the ChiSquare distribution. 
func ChiSquareExKurt(n int64) float64 {
	return float64(12/ n)
}

