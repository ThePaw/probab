// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Polya distribution. 
// Extension of the negative binomial distribution to the case of a positive real parameter r. 
//
// Parameters: 
// r > 0	 	number of failures until the experiment is stopped (integer, but the definition can also be extended to reals)
// p ∈ (0,1)	probability of success in each trial
//
// Support: 
// k ∈ { 0, 1, 2, 3, … }		number of successes

import (
	. "code.google.com/p/go-fn/fn"
	"math"
)

// PolyaPMF returns the PMF of the Polya distribution. 
func PolyaPMF(ρ, r float64) func(k int64) float64 {
	return func(k int64) float64 {
		kk := float64(k)
		return (Γ(kk+r) / (float64(Fact(k)) * Γ(r))) * math.Pow(1-ρ, r) * math.Pow(ρ, float64(k))
	}
}

// PolyaPMFAt returns the value of PMF of Polya distribution(μ, σ) at k. 
func PolyaPMFAt(ρ, r float64, k int64) float64 {
	pmf := PolyaPMF(ρ, r)
	return pmf(k)
}

// PolyaCDF returns the CDF of the Polya distribution. 
func PolyaCDF(ρ, r float64) func(k int64) float64 {
	return func(k int64) float64 {
		Ip := BetaCDFAt(float64(k+1), r, ρ)
		return 1 - Ip
	}
}

// PolyaCDFAt returns the value of CDF of the Polya distribution, at x. 
func PolyaCDFAt(ρ, r float64, k int64) float64 {
	cdf := PolyaCDF(ρ, r)
	return cdf(k)
}

// PolyaMean returns the mean of the Polya distribution. 
func PolyaMean(ρ, r float64) float64 {
	return ρ * r / (1 - ρ)
}

// PolyaMode returns the mode of the Polya distribution. 
func PolyaMode(ρ, r float64) float64 {
	if r > 1 {
		return math.Floor(ρ * (r - 1) / (1 - ρ))
	}
	return 0
}

// PolyaVar returns the variance of the Polya distribution. 
func PolyaVar(ρ, r float64) float64 {
	return ρ * r / ((1 - ρ) * (1 - ρ))
}

// PolyaStd returns the standard deviation of the Polya distribution. 
func PolyaStd(ρ, r float64) float64 {
	return math.Sqrt(ρ*r) / (1 - ρ)
}

// PolyaSkew returns the skewness of the Polya distribution. 
func PolyaSkew(ρ, r float64) float64 {
	return (1 + ρ) / math.Sqrt(ρ*r)
}

// PolyaExKurt returns the excess kurtosis of the Polya distribution. 
func PolyaExKurt(ρ, r float64) float64 {
	return 6/r + ((1-ρ)*(1-ρ))/(ρ*r)
}

// PolyaMGF returns the moment-generating function of the Polya distribution. 
func PolyaMGF(ρ, r float64, t float64) float64 {
	return math.Pow((1-ρ)/(1-ρ*math.Exp(t)), r)
}

// PolyaPGF returns the probability-generating function of the Polya distribution. 
func PolyaPGF(ρ, r float64, z float64) float64 {
	return math.Pow((1-ρ)/(1-ρ*z), r)
}
