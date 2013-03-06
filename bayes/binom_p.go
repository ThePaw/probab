// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Bayesian inference about the parameter p of binomial distribution.
// Bolstad 2007 (2e): Chapter 8, p. 141 and further.

import (
	. "code.google.com/p/probab/dst"
	"fmt"
	"math"
)

// Binomial proportion, posterior PDF, Flat prior.
func BinomPiPDFFPri(k, n int64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	α := 1.0
	β := 1.0
	return BetaPDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior PDF, Jeffreys prior.
// see Aitkin 2010: 143 for cautions
func BinomPiPDFJPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.5
	β = 0.5
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return BetaPDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior PDF, Haldane prior.
// see Aitkin 2010: 143 for cautions
func BinomPiPDFHPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.0
	β = 0.0
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return BetaPDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior PDF, general Beta prior.
func BinomPiPDFBPri(k, n int64, α, β float64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return BetaPDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior CDF, Flat prior.
func BinomPiCDFFPri(k, n int64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	α := 1.0
	β := 1.0
	return BetaCDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior CDF, Jeffreys prior.
// see Aitkin 2010: 143 for cautions
func BinomPiCDFJPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.5
	β = 0.5
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return BetaCDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior CDF, Haldane prior.
// see Aitkin 2010: 143 for cautions
func BinomPiCDFHPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.0
	β = 0.0
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return BetaCDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior CDF, general Beta prior.
func BinomPiCDFBPri(k, n int64, α, β float64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return BetaCDF(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior quantile function, Flat prior.
func BinomPiQtlFPri(k, n int64) func(p float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	α := 1.0
	β := 1.0
	return BetaQtl(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior quantile function, Jeffreys prior.
// see Aitkin 2010: 143 for cautions
func BinomPiQtlJPri(k, n int64) func(p float64) float64 {
	var α, β float64
	α = 0.5
	β = 0.5
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return BetaQtl(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior quantile function, Haldane prior.
// see Aitkin 2010: 143 for cautions
func BinomPiQtlHPri(k, n int64) func(p float64) float64 {
	var α, β float64
	α = 0.0
	β = 0.0
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return BetaQtl(α+float64(k), β+float64(n-k))
}

// Binomial proportion, posterior quantile function, general Beta prior.
func BinomPiQtlBPri(k, n int64, α, β float64) func(p float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return BetaQtl(α+float64(k), β+float64(n-k))
}

// Binomial proportion, Equivalent sample size of the prior
func BinomPiEqvSize(α, β float64) int64 {
	return int64(math.Floor(α + β + 1))
}

// Binomial proportion, Posterior modus
func BinomPiPostModus(α, β float64, n, k int64) float64 {
	var postα, postβ float64
	postα = α + float64(k)
	postβ = β + float64(n-k)
	return (postα - 1) / (postα + postβ - 2.0)
}

// Binomial proportion, Posterior mean
func BinomPiPostMean(α, β float64, n, k int64) float64 {
	var postα, postβ float64
	postα = α + float64(k)
	postβ = β + float64(n-k)
	return ((postα) / (postα + postβ))
}

// Binomial proportion, Posterior median
func BinomPiPostMedian(α, β float64, n, k int64) float64 {
	return 0 // to be implemented
}

// Binomial proportion, Posterior variance
// Bolstad 2007 (2e): 151, eq. 8.5
func BinomPiPostVar(α, β float64, n, k int64) float64 {
	var postα, postβ float64
	postα = α + float64(k)
	postβ = β + float64(n-k)
	return (postα * postβ) / ((postα + postβ) * (postα + postβ) * (postα + postβ + 1.0))
}

// Binomial proportion, Posterior mean square of p
// Bolstad 2007 (2e): 152-153, eq. 8.7
func BinomPiPMS(α, β float64, n, k, whichpi int64) float64 {
	const (
		MEAN   = 0
		MEDIAN = 1
		MODUS  = 2
	)

	var postmean, postvar, pihat float64

	postvar = BinomPiPostVar(α, β, n, k)
	postmean = BinomPiPostMean(α, β, n, k)

	switch whichpi {
	case MEAN:
		pihat = BinomPiPostMean(α, β, n, k)
	case MEDIAN:
		pihat = BinomPiPostMedian(α, β, n, k)
	case MODUS:
		pihat = BinomPiPostModus(α, β, n, k)
	}
	return postvar + (postmean-pihat)*(postmean-pihat)
}

// Binomial proportion, credible interval, beta prior, equal tail area
// Bolstad 2007 (2e): 153
// untested ...

func BinomPiCrIBP(α, β, alpha float64, n, k int64) (float64, float64) {
	/*
		k			observed successes
		n			total number of observations
		α			beta prior a
		β			beta prior b
		alpha			posterior probability that the true proportion lies outside the credible interval
	*/

	var low, upp float64
	low = BetaQtlFor(alpha/2.0, α+float64(k), β+float64(n-k))
	upp = BetaQtlFor(1.0-alpha/2.0, α+float64(k), β+float64(n-k))
	return low, upp
}

// Binomial proportion, credible interval, beta prior, equal tail area, normal approximation
// Bolstad 2007 (2e): 154-155, eq. 8.8
// untested ...

func BinomPiCrIBPriNApprox(α, β, alpha float64, n, k int64) (float64, float64) {
	/*
		k			observed successes
		n			total number of observations
		a			beta prior a
		b			beta prior b
		alpha			posterior probability that the true proportion lies outside the credible interval
	*/

	var postmean, postvar, postα, postβ, z, low, upp float64

	postα = α + float64(k)
	postβ = β + float64(n-k)

	postmean = postα / (postα + postβ)
	postvar = (postα * postβ) / ((postα + postβ) * (postα + postβ) * (postα + postβ + 1.0))
	z = ZQtlFor(alpha / 2)

	low = postmean - z*math.Sqrt(postvar)
	upp = postmean + z*math.Sqrt(postvar)
	return low, upp
}

// Binomial proportion, Likelihood
func BinomPiLike(pi float64, n, k int64) float64 {
	return math.Pow(pi, float64(k)) * math.Pow(1-pi, float64(n-k))
}

// Binomial proportion, Deviance 
func BinomPiDeviance(pi float64, n, k int64) float64 {
	return -2 * math.Log(BinomPiLike(pi, n, k))
}

// Binomial proportion, Sampling from posterior, Beta prior
func BinomPiCDFBPriNext(k, n int64, α, β float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return BetaNext(α+float64(k), β+float64(n-k))
}

// Binomial proportion, Deviance difference of a point null hypothesis pi = p against general alternative pi != p
// Aitkin 2010:143-144.
func binomPiPointDevDiff(k, n int64, α, β, p, pi float64) float64 {
	nn := float64(n)
	kk := float64(k)
	d0 := -2 * (kk*math.Log(p) + (nn-kk)*math.Log(1-p)) //  null model deviance
	dd := d0 + 2*(kk*math.Log(pi)+(nn-kk)*math.Log(1-pi))
	return dd
}
