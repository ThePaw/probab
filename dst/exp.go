// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Exponential distribution  (a.k.a. negative exponential distribution). 
// Parameters:
// λ > 0: rate, or inverse scale
// Support: x ∈ [0; ∞).


import (
	"math"
	"math/rand"
)

// Exponential_PDF returns the PDF of the Exponential distribution. 
func Exp_PDF(λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return λ * math.Exp(-λ*x)
	}
}


// Exponential_LnPDF returns the natural logarithm of the PDF of the Exponential distribution. 
func Exp_LnPDF(λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return negInf
		}
		return log(λ) - λ*x
	}
}

// Exponential_PDF_At returns the value of PDF of Exponential distribution at x. 
func Exponential_PDF_At(λ, x float64) float64 {
	pdf := Exp_PDF(λ)
	return pdf(x)
}

// Exponential_CDF returns the CDF of the Exponential distribution. 
func Exponential_CDF(λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return 1- math.Exp(-λ*x)
	}
}

// Exponential_CDF_At returns the value of CDF of the Exponential distribution, at x. 
func Exponential_CDF_At(λ, x float64) float64 {
	cdf := Exponential_CDF(λ)
	return cdf(x)
}

// Exponential_Qtl returns the inverse of the CDF (quantile) of the Exponential distribution. 
func Exponential_Qtl(λ float64) func(p float64) float64 {
	// p: probability for which the quantile is evaluated
	return func(p float64) float64 {
		return -math.Log(1-p)/λ
	}
}

// Exponential_Qtl_For returns the inverse of the CDF (quantile) of the Exponential distribution, for given probability.
func Exponential_Qtl_For(λ, p float64) float64 {
	cdf := Exponential_Qtl(λ)
	return cdf(p)
}

// NextExponential returns random number drawn from the Exponential distribution. 
func NextExp(λ float64) float64    { return rand.ExpFloat64() / λ }

// Exponential returns the random number generator with  Exponential distribution. 
func Exp(λ float64) func() float64 { return func() float64 { return NextExp(λ) } }

// ExponentialMean returns the mean of the Exponential distribution. 
func ExponentialMean(λ float64) float64 {
	return 1/λ
}

// ExponentialMedian returns the median of the Exponential distribution. 
func ExponentialMedian(λ float64) (med float64) {
	return (1/λ)*math.Log(2)
}

// ExponentialMode returns the mode of the Exponential distribution. 
func ExponentialMode(λ float64) float64 {
	return 0
}

// ExponentialVar returns the variance of the Exponential distribution. 
func ExponentialVar(λ float64) float64 {
	return 1/(λ*λ)
}

// ExponentialStd returns the standard deviation of the Exponential distribution. 
func ExponentialStd(λ float64) float64 {
	return 1/λ
}

// ExponentialSkew returns the skewness of the Exponential distribution. 
func ExponentialSkew(λ float64) (s float64) {
	return 2
}

// ExponentialExKurt returns the excess kurtosis of the Exponential distribution. 
func ExponentialExKurt(λ float64) float64 {
	return 6
}

// ExponentialMGF returns the moment-generating function of the Exponential distribution. 
func ExponentialMGF(λ, p, t float64) float64 {
	return 1/(1-t/λ)
}
