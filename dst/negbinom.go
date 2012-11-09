// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Negative binomial distribution. 
// A discrete probability distribution of the number of successes in a sequence of Bernoulli trials before a specified (non-random) number of failures (denoted r) occur. For example, if one throws a die repeatedly until the third time “1” appears, then the probability distribution of the number of non-“1”s that had appeared will be negative binomial.
//
// Parameters: 
// r > 0	 	number of failures until the experiment is stopped (integer, but the definition can also be extended to reals)
// p ∈ (0,1)	probability of success in each trial
//
// Support: 
// k ∈ { 0, 1, 2, 3, … }		number of successes

import (
	"math"
	. "code.google.com/p/go-fn/fn"
)

// NegBinomial_PMF returns the PMF of the NegBinomial distribution. 
func NegBinomial_PMF(ρ float64, r int64) func(k int64) float64 {
	return func(k int64) float64 {
		return BinomCoeff(k + r - 1, k) * math.Pow(1-ρ, float64(r)) * math.Pow(ρ, float64(k))
	}
}

// NegBinomial_LnPMF returns the natural logarithm of the PMF of the NegBinomial distribution. 
func NegBinomial_LnPMF(ρ float64, r int64) func(i int64) float64 {
	return func(k int64) float64 {
		rr := float64(r)
		return LnChoose(k+r-1, r-1) + log(ρ)*rr + log(1-ρ)*float64(k)
	}
}

// NegBinomial_PMF_At returns the value of PMF of NegBinomial distribution(μ, σ) at k. 
func NegBinomial_PMF_At(ρ float64, r, k int64)  float64 {
	pmf := NegBinomial_PMF(ρ, r) 
	return pmf(k)
}

// NegBinomial_CDF returns the CDF of the NegBinomial distribution. 
func NegBinomial_CDF(ρ float64, r int64) func(k int64) float64 {
	return func(k int64) float64 {
		Ip:=Beta_CDF_At(float64(k+1), float64(r), ρ)
		return 1-Ip
	}
}

// NegBinomial_CDF_At returns the value of CDF of the NegBinomial distribution, at x. 
func NegBinomial_CDF_At(ρ float64, r, k int64)  float64 {
	cdf := NegBinomial_CDF(ρ, r) 
	return cdf(k)
}

// NextNegBinomial returns random number drawn from the NegBinomial distribution. 
func NextNegBinomial(ρ float64, r int64) int64 {
	k := iZero
	for r >= 0 {
		i := NextBernoulli(ρ)
		r -= i
		k += (1 - i)
	}
	return k
}

// NegBinomial returns the random number generator with  NegBinomial distribution. 
func NegBinomial(ρ float64, r int64) func() int64 {
	return func() int64 {
		return NextNegBinomial(ρ, r)
	}
}

// NegBinomialMean returns the mean of the NegBinomial distribution. 
func NegBinomialMean(ρ float64, r int64) float64 {
	return ρ*float64(r)/(1-ρ)
}

// NegBinomialMode returns the mode of the NegBinomial distribution. 
func NegBinomialMode(ρ float64, r int64) float64 {
	if r > 1  {
		return math.Floor(ρ*float64(r-1)/(1-ρ))
	}
	return 0
}

// NegBinomialVar returns the variance of the NegBinomial distribution. 
func NegBinomialVar(ρ float64, r int64) float64 {
	return ρ*float64(r)/((1-ρ)*(1-ρ))
}

// NegBinomialStd returns the standard deviation of the NegBinomial distribution. 
func NegBinomialStd(ρ float64, r int64) float64 {
	return math.Sqrt(ρ*float64(r))/(1-ρ)
}

// NegBinomialSkew returns the skewness of the NegBinomial distribution. 
func NegBinomialSkew(ρ float64, r int64) float64 {
	return (1+ρ)/math.Sqrt(ρ*float64(r))
}

// NegBinomialExKurt returns the excess kurtosis of the NegBinomial distribution. 
func NegBinomialExKurt(ρ float64, r int64) float64 {
	rr := float64(r)
	return 6/rr + ((1-ρ)*(1-ρ))/(ρ*rr)
}

// NegBinomialMGF returns the moment-generating function of the NegBinomial distribution. 
func NegBinomialMGF(ρ float64, r int64, t float64) float64 {
	return math.Pow((1-ρ)/ (1-ρ*math.Exp(t)), float64(r))
}

// NegBinomialPGF returns the probability-generating function of the NegBinomial distribution. 
func NegBinomialPGF(ρ float64, r int64, z float64) float64 {
	return math.Pow((1-ρ)/ (1-ρ*z), float64(r))
}

