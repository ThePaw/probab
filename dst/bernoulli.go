// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Bernoulli distribution.

// Bernoulli_PMF returns the PMF of the Bernoulli distribution. 
func Bernoulli_PMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k < 0 || k > 1 {
			panic("k is not 0 or 1")
		}
		if k == 1 {
			return ρ
		}
		return 1 - ρ
	}
}

// Bernoulli_LnPMF returns the natural logarithm of the PMF of the Bernoulli distribution. 
func Bernoulli_LnPMF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k == 1 {
			return log(ρ)
		}
		return log(1 - ρ)
	}
}

// Bernoulli_PMF_At returns the value of PMF of Bernoulli distribution at x. 
func Bernoulli_PMF_At(ρ float64, k int64) float64 {
	pmf := Bernoulli_PMF(ρ)
	return pmf(k)
}

// Bernoulli_CDF returns the value of CDF of the Bernoulli distribution, at x. 
func Bernoulli_CDF(ρ float64) func(k int64) float64 {
	return func(k int64) float64 {
		if k < 0 || k > 1 {
			panic("k is not 0 or 1")
		}
		if k == 1 {
			return 1
		}
		return 1 - ρ
	}
}

// Bernoulli_CDF_At returns the value of CDF of the Bernoulli distribution, at x. 
func Bernoulli_CDF_At(ρ float64, k int64)  float64 {
	cdf := Bernoulli_CDF(ρ)
	return cdf(k)
}

// NextBernoulli returns random number drawn from the Bernoulli distribution. 
func NextBernoulli(ρ float64) int64 {
	if NextUniform(0, 1) < ρ {
		return 1
	}
	return 0
}

// Bernoulli returns the random number generator with  Bernoulli distribution. 
func Bernoulli(ρ float64) func() int64 { return func() int64 { return NextBernoulli(ρ) } }

