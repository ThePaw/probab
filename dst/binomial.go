// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Binomial distribution. 
// Parameters: 
// n ∈ N0	 	number of trials
// p ∈ [0, 1]	probability of success in each trial
// Support: 
// k ∈ {0, ... , n}

import (
	. "code.google.com/p/go-fn/fn"
	"math"
)

// BinomialPMF returns the PMF of the Binomial distribution. 
func BinomialPMF(n int64, p float64) func(k int64) float64 {
	return func(k int64) (x float64) {
		x = pow(p, float64(k)) * pow(1-p, float64(n-k))
		x *= Γ(float64(n+1)) / (Γ(float64(k+1)) * Γ(float64(n-k+1)))
		return
	}
}

// BinomialLnPMF returns the natural logarithm of the PMF of the Binomial distribution. 
func BinomialLnPMF(n int64, p float64) func(k int64) float64 {
	return func(k int64) (x float64) {
		x = log(p)*float64(k) + log(1-p)*float64(n-k)
		x += LnΓ(float64(n+1)) - LnΓ(float64(k+1)) - LnΓ(float64(n-k+1))
		return
	}
}

// BinomialPMFAt returns the value of PMF of Binomial distribution at k. 
func BinomialPMFAt(n int64, p float64, k int64) float64 {
	pmf := BinomialPMF(n, p)
	return pmf(k)
}

// BinomialCDF returns the CDF of the Binomial distribution. 
func BinomialCDF(n int64, p float64) func(k int64) float64 {
	return func(k int64) float64 {
		return BetaCDFAt((float64)(n-k), (float64)(k+1), 1-p)
	}
}

// BinomialCDFAt returns the value of CDF of the Binomial distribution, at k. 
func BinomialCDFAt(n int64, p float64, k int64) float64 {
	cdf := BinomialCDF(n, p)
	return cdf(k)
}

// BinomialQtl returns the inverse of the CDF (quantile) of the Binomial distribution.
func BinomialQtl(n int64, ρ float64) func(p float64) int64 {
	return func(p float64) int64 {
		var eps, q, mu, sigma, gamma, z  float64
		var y int64
		eps = 2.2204460492503131e-16 // DBL_EPSILON

		if float64(n) != math.Floor(float64(n)+0.5) {
			panic("bad n")
		}
		if ρ < 0 || ρ > 1 || n < 0 {
			panic("bad params")
		}

		if ρ == 0 || n == 0 {
			return 0
		}

		q = 1 - ρ
		if q == 0 {
			return n // covers the full range of the distribution 
		}
		mu = float64(n) * ρ
		sigma = sqrt(float64(n) * ρ * q)
		gamma = (q - ρ) / sigma

		// temporary hack --- FIXME ---
		if p+1.01*eps >= 1 {
			return n
		}

		// y = apρox.value (Cornish-Fisher expansion)
		z = NormalQtlFor(0, 1, p)
		y = int64(math.Floor(mu + sigma*(z+gamma*(z*z-1)/6) + 0.5))

		if y > n { // way off  
			y = n
		}
		z = BinomialCDFAt(n, ρ, y)
		// fuzz to ensure left continuity
		p *= 1 - 64*eps

		// If the C-F value is not too large a simple search is OK
		if y < 1e5 {
			return searchBinomial(p, ρ, y, n, 1, &z)
		}
		// Otherwise be a bit cleverer in the search 
		{
			incr := int64(math.Floor(float64(n) / 1000))
			oldincr := incr
			for oldincr > 1 && incr > int64(math.Floor(float64(y)*1e-15)) {
				y = searchBinomial(p, ρ, y, n, incr, &z)
				incr = imax(1, incr/100)
				oldincr = incr
			}
			return y
		}
	}
}

// BinomialQtlFor returns the inverse of the CDF (quantile) of the Negative binomial distribution, for given probability.
func BinomialQtlFor(n int64, ρ, p float64) int64 {
	qtl := BinomialQtl(n, ρ)
	return qtl(p)
}

// BinomialNext returns random number drawn from the Binomial distribution. 
func BinomialNext(n int64, p float64) (x int64) {
	x = 0
	for i := int64(0); i <= n; i++ {
		x += BernoulliNext(p)
	}
	return
}

// Binomial returns the random number generator with  Binomial distribution. 
func Binomial(n int64, p float64) func() int64 {
	return func() int64 { return BinomialNext(n, p) }
}

// BinomialMean returns the mean of the Binomial distribution. 
func BinomialMean(n int64, p float64) float64 {
	return float64(n) * p
}

// BinomialMedian returns the median of the Binomial distribution. 
func BinomialMedian(n int64, p float64) float64 {
	return math.Floor(float64(n) * p)
}

// BinomialMode returns the mode of the Binomial distribution. 
func BinomialMode(n int64, p float64) float64 {
	ε := 1e-3 // some small number
	switch {
	case (float64(n+1)*p)-math.Floor(float64(n+1)*p) > ε: // (n+1)*p is non-integer
		return math.Floor(float64(n+1) * p)
	case (float64(n+1) * p) <= ε: // (n+1)*p == 0
		return math.Floor(float64(n+1) * p)
	case (float64(n+1)*p)-float64(n+1) <= ε: // (n+1)*p == (n+1)
		return float64(n)
	}
	return float64(n+1) * p // (n+1)*p is integer
}

// BinomialVar returns the variance of the Binomial distribution. 
func BinomialVar(n int64, p float64) float64 {
	return float64(n) * p * (1 - p)
}

// BinomialStd returns the standard deviation of the Binomial distribution. 
func BinomialStd(n int64, p float64) float64 {
	return math.Sqrt(float64(n) * p * (1 - p))
}

// BinomialSkew returns the skewness of the Binomial distribution. 
func BinomialSkew(n int64, p float64) float64 {
	return 1 - 2*p/math.Sqrt(float64(n)*p*(1-p))
}

// BinomialExKurt returns the excess kurtosis of the Binomial distribution. 
func BinomialExKurt(n int64, p float64) float64 {
	return 1 - 6*p*(1-p)/(float64(n)*p*(1-p))
}

// BinomialMGF returns the moment-generating function of the Binomial distribution. 
func BinomialMGF(n int64, p, t float64) float64 {
	return math.Pow((1 - p + p*math.Exp(t)), float64(n))
}

// BinomialPGF returns the probability-generating function of the Binomial distribution. 
func BinomialPGF(n int64, p, z float64) float64 {
	return math.Pow((1 - p + p*z), float64(n))
}

func searchBinomial(p, pr float64, y, n, incr int64, z *float64) int64 {
	if *z >= p {
		// search to the left
		for {
			newz := BinomialCDFAt(n, pr,  y-incr)
			if y == 0 || newz < p {
				return y
			}
			y = imax(0, y-incr)
			*z = newz
		}
	} else { // search to the right
		for {
	    		y = imin(y + incr, n)
			*z = BinomialCDFAt(n, pr,  y)
			if y == n || *z >= p {
				return y
			}
		}
	}
	return y // just to make compiler happy ;-)
}

