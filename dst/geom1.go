// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Geometric distribution (type 1). 
// The probability distribution of the number Y = X − 1 of failures before the first success, supported on the set {1, 2, 3, ... }
// Parameters: 
// ρ ∈ (0, 1]	probability of success in each trial
// Support: 
// k ∈ {1, ... , n}

import "math"

// Geometric1_PMF returns the PMF of the Geometric1 distribution. 
func Geometric1_PMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 { return ρ * pow(1-ρ, float64(k-1)) }
}

// Geometric1_LnPMF returns the natural logarithm of the PMF of the Geometric1 distribution. 
func Geometric1_LnPMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 { return log(1-ρ) + float64(k-1)*log(ρ) }
}

// Geometric1_PMF_At returns the value of PMF of Geometric1 distribution at k. 
func Geometric1_PMF_At(ρ float64, k int64) float64 {
	pmf := Geometric1_PMF(ρ)
	return pmf(k)
}

// Geometric1_CDF returns the value of CDF of the Geometric1 distribution, at k. 
func Geometric1_CDF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k <  0 {
			panic("k < 0")
		}
		return 1 - pow(1-ρ, float64(k)) 
	}
}

// Geometric1_CDF_At returns the value of CDF of the Geometric1 distribution, at x. 
func Geometric1_CDF_At(ρ float64, k int64)  float64 {
	cdf := Geometric1_CDF(ρ)
	return cdf(k)
}

/* Not tested, looking strange, commented out, waiting for revision
// NextGeometric1 returns random number drawn from the Geometric1 distribution. 
//NextGeometric1(ρ) => # of NextGeometric1(ρ) failures before one success
func NextGeometric1(ρ float64) int64 {
	if NextGeometric1(ρ) == 1 {
		return 1 + NextGeometric1(ρ)
	}
	return 0
}

// Geometric1 returns the random number generator with  Geometric1 distribution. 
func Geometric1(ρ float64) func() int64 { return func() int64 { return NextGeometric1(ρ) } }
*/

// Geometric1Mean returns the mean of the Geometric1 distribution. 
func Geometric1Mean(ρ float64) float64 {
	return 1/ρ
}

/*  to be implemented
// Geometric1Median returns the median of the Geometric1 distribution. 
func Geometric1Median(ρ float64) float64 {
	return math.Floor(float64(n)*p)
}
*/

// Geometric1Mode returns the mode of the Geometric1 distribution. 
func Geometric1Mode(ρ float64) float64 {
	return 1
}

// Geometric1Var returns the variance of the Geometric1 distribution. 
func Geometric1Var(ρ float64) float64 {
	return (1-ρ)/(ρ*ρ)
}

// Geometric1Std returns the standard deviation of the Geometric1 distribution. 
func Geometric1Std(ρ float64) float64 {
	return math.Sqrt(1-ρ) / ρ
}

// Geometric1Skew returns the skewness of the Geometric1 distribution. 
func Geometric1Skew(ρ float64) float64 {
	return (2-ρ) / math.Sqrt(1-ρ)
}

// Geometric1ExKurt returns the excess kurtosis of the Geometric1 distribution. 
func Geometric1ExKurt(ρ float64) float64 {
	return 6 + (ρ*ρ)/(1-ρ)
} 

// Geometric1MGF returns the moment-generating function of the Geometric1 distribution. 
func Geometric1MGF(ρ, t float64) float64 {
	if t >= -math.Log(1-ρ) {
		panic("MGF not defined for t >= -math.Log(1-ρ)")
	}
	return ρ*math.Exp(t)/(1-(1-ρ)*math.Exp(t))
}
