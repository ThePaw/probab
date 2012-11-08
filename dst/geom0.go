// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Geometric distribution (type 0). 
// The probability distribution of the number Y = X − 1 of failures before the first success, supported on the set { 0, 1, 2, 3, ... }
// Parameters: 
// ρ ∈ (0, 1]	probability of success in each trial
// Support: 
// k ∈ {0, ... , n}

import "math"

// Geometric_PMF returns the PMF of the Geometric distribution. 
func Geometric_PMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 { return ρ * pow(1-ρ, float64(k)) }
}

// Geometric_LnPMF returns the natural logarithm of the PMF of the Geometric distribution. 
func Geometric_LnPMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 { return log(1-ρ) + float64(k)*log(ρ) }
}

// Geometric_PMF_At returns the value of PMF of Geometric distribution at k. 
func Geometric_PMF_At(ρ float64, k int64) float64 {
	pmf := Geometric_PMF(ρ)
	return pmf(k)
}

// Geometric_CDF returns the value of CDF of the Geometric distribution, at k. 
func Geometric_CDF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k <  0 {
			panic("k < 0")
		}
		return 1 - pow(1-ρ, float64(k+1)) 
	}
}

// Geometric_CDF_At returns the value of CDF of the Geometric distribution, at x. 
func Geometric_CDF_At(ρ float64, k int64)  float64 {
	cdf := Geometric_CDF(ρ)
	return cdf(k)
}

/* Not tested, looking strange, commented out, waiting for revision
// NextGeometric returns random number drawn from the Geometric distribution. 
//NextGeometric(ρ) => # of NextGeometric(ρ) failures before one success
func NextGeometric(ρ float64) int64 {
	if NextGeometric(ρ) == 1 {
		return 1 + NextGeometric(ρ)
	}
	return 0
}

// Geometric returns the random number generator with  Geometric distribution. 
func Geometric(ρ float64) func() int64 { return func() int64 { return NextGeometric(ρ) } }
*/

// GeometricMean returns the mean of the Geometric distribution. 
func GeometricMean(ρ float64) float64 {
	return (1-ρ)/ρ
}

/*  to be implemented
// GeometricMedian returns the median of the Geometric distribution. 
func GeometricMedian(ρ float64) float64 {
	return math.Floor(float64(n)*p)
}
*/

// GeometricMode returns the mode of the Geometric distribution. 
func GeometricMode(ρ float64) float64 {
	return 0
}

// GeometricVar returns the variance of the Geometric distribution. 
func GeometricVar(ρ float64) float64 {
	return (1-ρ)/(ρ*ρ)
}

// GeometricStd returns the standard deviation of the Geometric distribution. 
func GeometricStd(ρ float64) float64 {
	return math.Sqrt(1-ρ) / ρ
}

// GeometricSkew returns the skewness of the Geometric distribution. 
func GeometricSkew(ρ float64) float64 {
	return (2-ρ) / math.Sqrt(1-ρ)
}

// GeometricExKurt returns the excess kurtosis of the Geometric distribution. 
func GeometricExKurt(ρ float64) float64 {
	return 6 + (ρ*ρ)/(1-ρ)
} 

// GeometricMGF returns the moment-generating function of the Geometric distribution. 
func GeometricMGF(ρ, t float64) float64 {
	return ρ/(1-(1-ρ)*math.Exp(t))
}
