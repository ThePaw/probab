// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Hypergeometric distribution. 
// A discrete probability distribution that describes the probability of k successes in n draws 
// from a finite population of size nN containing m successes without replacement. 
// Parameters: 
// nN ∈ {1, 2,  ... }
// m ∈ {0, 1, 2,  ... , nN}
// n ∈ {1, 2,  ... , nN}
// Support: 
// k ∈ {max(0, n+m-nN), ... , min(m, n)}

// HypergeometricPMF returns the PMF of the Hypergeometric distribution. 
func HypergeometricPMF(nN, m, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		if nN < 1 || m < 0 || m > nN || n < 1 || n > nN {
			return NaN
		}
		// p := BinomCoeff(m, k) * BinomCoeff(nN-m, n-k)  / BinomCoeff(nN, n) 
		return exp(logBinomCoeff(m, k) + logBinomCoeff(nN-m, n-k) - logBinomCoeff(nN, n))
	}
}

// HypergeometricLnPMF returns the natural logarithm of the PMF of the Hypergeometric distribution. 
func HypergeometricLnPMF(nN, m, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		if nN < 1 || m < 0 || m > nN || n < 1 || n > nN {
			return NaN
		}
		return logBinomCoeff(m, k) + logBinomCoeff(nN-m, n-k) - logBinomCoeff(nN, n)
	}
}

// HypergeometricPMFAt returns the value of PMF of Hypergeometric distribution(μ, σ) at k. 
func HypergeometricPMFAt(nN, m, n, k int64) float64 {
	if float64(k) < max(0, float64(n+m-nN)) || float64(k) > min(float64(m), float64(n)) {
		return NaN
	}
	pmf := HypergeometricPMF(nN, m, n)
	return pmf(k)
}

// HypergeometricCDF returns the CDF of the Hypergeometric distribution. 
func HypergeometricCDF(nN, m, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		var (
			p float64 = 0.0
			i int64
		)
		pmf := HypergeometricPMF(nN, m, n)
		for i = 0; i <= k; i++ {
			p += pmf(i)
		}
		return p
	}
}

// HypergeometricCDFAt returns the value of CDF of the Hypergeometric distribution, at k.
func HypergeometricCDFAt(nN, m, n, k int64) float64 {
	cdf := HypergeometricCDF(nN, m, n)
	return cdf(k)
}

//		=== Approximations using standard normal distribution function ===
//		Only use iff n is large, nN and m are large compared to n 
//		and p = m/nN is not close to 0 or 1

// HypergeometricApproxPMF returns the PMF of the Hypergeometric distribution approximated using Standard normal distribution. 
func HypergeometricApproxPMF(nN, m, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		if nN < 1 || m < 0 || m > nN || n < 1 || n > nN {
			return NaN
		}
		p := float64(m) / float64(nN)
		x := float64(k-n) * p / sqrt(float64(n)*p*(1-p))
		return ZPDFAt(x)
	}
}

// HypergeometricApproxPMFAt returns the value of PMF of Hypergeometric distribution approximated using Standard normal distribution, at k. 
func HypergeometricApproxPMFAt(nN, m, n, k int64) float64 {
	if float64(k) < max(0, float64(n+m-nN)) || float64(k) > min(float64(m), float64(n)) {
		return NaN
	}
	pmf := HypergeometricApproxPMF(nN, m, n)
	return pmf(k)
}

// HypergeometricApproxCDF returns the CDF of the Hypergeometric distribution approximated using Standard normal distribution. 
func HypergeometricApproxCDF(nN, m, n int64) func(k int64) float64 {
	return func(k int64) float64 {
		p := float64(m) / float64(nN)
		x := float64(k-n) * p / sqrt(float64(n)*p*(1-p))
		return ZCDFAt(x)
	}
}

// HypergeometricApproxCDFAt returns the value of CDF of the Hypergeometric distribution approximated using Standard normal distribution, at k.
func HypergeometricApproxCDFAt(nN, m, n, k int64) float64 {
	cdf := HypergeometricApproxCDF(nN, m, n)
	return cdf(k)
}

//		=== 

// HypergeometricMean returns the mean of the Hypergeometric distribution. 
func HypergeometricMean(nN, m, n int64) float64 {
	return float64(n*m) / float64(nN)
}

// HypergeometricMedian returns the median of the Hypergeometric distribution. 
// func HypergeometricMedian(nN, m, n int64) float64 {

// HypergeometricMode returns the mode of the Hypergeometric distribution. 
func HypergeometricMode(nN, m, n int64) float64 {
	return floor(float64(n+1) * float64(m+1) / float64(nN+2))
}

// HypergeometricVar returns the variance of the Hypergeometric distribution. 
func HypergeometricVar(nN, m, n int64) float64 {
	return float64(n) * (float64(m) / float64(nN)) * (float64(nN-m) / float64(nN)) * ((float64(nN - n)) / float64(nN-1))
}

// HypergeometricStd returns the standard deviation of the Hypergeometric distribution. 
func HypergeometricStd(nN, m, n int64) float64 {
	return sqrt(float64(n) * (float64(m) / float64(nN)) * (float64(nN-m) / float64(nN)) * ((float64(nN - n)) / float64(nN-1)))
}

// HypergeometricSkew returns the skewness of the Hypergeometric distribution. 
func HypergeometricSkew(nN, m, n int64) float64 {
	num := float64(nN-2*m) * sqrt(float64(nN-1)) * float64(nN-2*n)
	den := sqrt(float64(n*m)*float64(nN-m)*float64(nN-n)) * float64(nN-2)
	return num / den
}

// HypergeometricExKurt returns the excess kurtosis of the Hypergeometric distribution. 
func HypergeometricExKurt(nN, m, n int64) float64 {
	num := ((nN-1)*nN*nN*(nN*(nN+1)-6*m*(nN-m)-6*n*(nN-n)) + 6*n*m*(nN-m)*(nN-n)*(5*nN-6))
	//	den := float64(n*m)*float64(nN-m) * float64(nN-n)* float64(nN-2)* float64(nN-3)
	den := n * m * (nN - m) * (nN - n) * (nN - 2) * (nN - 3)
	return float64(num) / float64(den)
}

/* To be implemented ...
// HypergeometricMGF returns the moment-generating function of the Hypergeometric distribution. 
func HypergeometricMGF(n int64, p, t float64) float64 {
	return 
}
*/
