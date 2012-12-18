// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Inverse-gamma distribution (not to be confused with Inverse CDF of Gamma distribution). 
//
// A two-parameter family of continuous probability distributions on the positive real line, 
// which is the distribution of the reciprocal of a variable distributed according to the gamma distribution. 
// Perhaps the chief use of the inverse gamma distribution is in Bayesian statistics, 
// where it serves as the conjugate prior of the variance of a normal distribution. 
// However, it is common among Bayesians to consider an alternative parametrization 
// of the normal distribution in terms of the precision, defined as the reciprocal of the variance, 
// which allows the gamma distribution to be used directly as a conjugate prior.
//
// Parameters: 
// α > 0:		shape
// β > 0:		scale
// Support:	x ∈ (0, ∞)

import (
	. "code.google.com/p/go-fn/fn"
	"math"
)

// InvGammaPDF returns the PDF of the InvGamma distribution. 
func InvGammaPDF(α, β float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Exp(α*math.Log(β) - LnΓ(α) - (α+1)*math.Log(x) - β*1.0/x)
	}
}

// InvGammaLnPDF returns the natural logarithm of the PDF of the InvGamma distribution. 
func InvGammaLnPDF(α, β float64) func(x float64) float64 {
	return func(x float64) float64 {
		return α*math.Log(β) - LnΓ(α) - (α+1)*math.Log(x) - β*1.0/x
	}
}

// InvGammaPDFAt returns the value of PDF of InvGamma distribution at x. 
func InvGammaPDFAt(α, β float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Exp(α*math.Log(β) - LnΓ(α) - (α+1)*math.Log(x) - β/x)
	}
}

// InvGammaCDF returns the CDF of the InvGamma distribution. 
func InvGammaCDF(α, β float64) func(x float64) float64 {
	return func(x float64) float64 {
		return 1 - IΓ(α, β/x)
	}
}

// InvGammaCDFAt returns the value of CDF of the InvGamma distribution, at x. 
func InvGammaCDFAt(α, β, x float64) float64 {
	cdf := InvGammaCDF(α, β)
	return cdf(x)
}

// InvGammaMean returns the mean of the InvGamma distribution. 
func InvGammaMean(α, β float64) float64 {
	if α <= 1 {
		panic("mean not defined for α <= 1")
	}
	return β / (α - 1)
}

// InvGammaMedian returns the median of the InvGamma distribution. 
// to be implemented ...

// InvGammaMode returns the mode of the InvGamma distribution. 
func InvGammaMode(α, β float64) float64 {
	return β / (α + 1)
}

// InvGammaVar returns the variance of the InvGamma distribution. 
func InvGammaVar(α, β float64) float64 {
	if α <= 2 {
		panic("variance not defined for α <=2 ")
	}
	return (β * β) / ((α - 1) * (α - 1) * (α - 2))
}

// InvGammaStd returns the standard deviation of the InvGamma distribution. 
func InvGammaStd(α, β float64) float64 {
	if α <= 2 {
		panic("standard deviation not defined for α <= 2")
	}
	return β / math.Sqrt((α-1)*(α-1)*(α-2))
}

// InvGammaSkew returns the skewness of the InvGamma distribution. 
func InvGammaSkew(α, β float64) (s float64) {
	if α <= 3 {
		panic("skewness not defined for α <= 3")
	}
	return 4 * math.Sqrt(α-2) / (α - 3)
}

// InvGammaExKurt returns the excess kurtosis of the InvGamma distribution. 
func InvGammaExKurt(α, β float64) float64 {
	if α <= 4 {
		panic("excess kurtosis not defined for α <= 3")
	}
	return (30*α - 66) / ((α - 3) * (α - 4))
}

// InvGammaMGF returns the moment-generating function of the InvGamma distribution. To be implemented ...
