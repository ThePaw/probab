// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Bayesian inference about the parameter p of binomial distribution.
// Bolstad 2007 (2e): Chapter 8, p. 141 and further.

import (
	"code.google.com/p/probab/dst"
	"fmt"
	"math"
)

// BinomPiPDFFPri returns posterior PDF of the Binomial proportion, Flat prior.
func BinomPiPDFFPri(k, n int64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	α := 1.0
	β := 1.0
	return dst.BetaPDF(α+float64(k), β+float64(n-k))
}

// BinomPiPDFJPri returns posterior PDFof the Binomial proportion, Jeffreys prior.
// see Aitkin 2010: 143 for cautions
func BinomPiPDFJPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.5
	β = 0.5
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return dst.BetaPDF(α+float64(k), β+float64(n-k))
}

// BinomPiPDFHPri returns posterior PDF of the Binomial proportion, Haldane prior.
// see Aitkin 2010: 143 for cautions
func BinomPiPDFHPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.0
	β = 0.0
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return dst.BetaPDF(α+float64(k), β+float64(n-k))
}

// BinomPiPDFBPri returns posterior PDF of the Binomial proportion, general Beta prior.
func BinomPiPDFBPri(k, n int64, α, β float64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return dst.BetaPDF(α+float64(k), β+float64(n-k))
}

// BinomPiCDFFPri returns posterior CDF of the Binomial proportion, Flat prior.
func BinomPiCDFFPri(k, n int64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	α := 1.0
	β := 1.0
	return dst.BetaCDF(α+float64(k), β+float64(n-k))
}

// BinomPiCDFJPri returns posterior CDF of the Binomial proportion, Jeffreys prior.
// see Aitkin 2010: 143 for cautions
func BinomPiCDFJPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.5
	β = 0.5
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return dst.BetaCDF(α+float64(k), β+float64(n-k))
}

// BinomPiCDFHPri returns posterior CDF of the Binomial proportion, Haldane prior.
// see Aitkin 2010: 143 for cautions
func BinomPiCDFHPri(k, n int64) func(x float64) float64 {
	var α, β float64
	α = 0.0
	β = 0.0
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return dst.BetaCDF(α+float64(k), β+float64(n-k))
}

// BinomPiCDFBPri returns posterior CDF of the Binomial proportion, general Beta prior.
func BinomPiCDFBPri(k, n int64, α, β float64) func(x float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return dst.BetaCDF(α+float64(k), β+float64(n-k))
}

// BinomPiQtlFPri returns posterior quantile function for Binomial proportion, Flat prior.
func BinomPiQtlFPri(k, n int64) func(p float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	α := 1.0
	β := 1.0
	return dst.BetaQtl(α+float64(k), β+float64(n-k))
}

// BinomPiQtlJPri returns  posterior quantile function for Binomial proportion, Jeffreys prior.
// see Aitkin 2010: 143 for cautions
func BinomPiQtlJPri(k, n int64) func(p float64) float64 {
	var α, β float64
	α = 0.5
	β = 0.5
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return dst.BetaQtl(α+float64(k), β+float64(n-k))
}

// BinomPiQtlHPri returns posterior quantile function for Binomial proportion, Haldane prior.
// see Aitkin 2010: 143 for cautions
func BinomPiQtlHPri(k, n int64) func(p float64) float64 {
	var α, β float64
	α = 0.0
	β = 0.0
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	return dst.BetaQtl(α+float64(k), β+float64(n-k))
}

// BinomPiQtlBPri returns posterior quantile function forBinomial proportion, general Beta prior.
func BinomPiQtlBPri(k, n int64, α, β float64) func(p float64) float64 {
	if k > n {
		panic(fmt.Sprintf("The number of observed successes (k) must be <= number of trials (n)"))
	}
	if α < 0 || β < 0 {
		panic(fmt.Sprintf("The parameters of the prior must be non-negative"))
	}
	return dst.BetaQtl(α+float64(k), β+float64(n-k))
}

// BinomPiEqvSize returns the Equivalent sample size of the prior of the Binomial proportion.
func BinomPiEqvSize(α, β float64) int64 {
	return int64(math.Floor(α + β + 1))
}

// BinomPiPostModus returns Posterior modus of the Binomial proportion.
func BinomPiPostModus(α, β float64, n, k int64) float64 {
	var postα, postβ float64
	postα = α + float64(k)
	postβ = β + float64(n-k)
	return (postα - 1) / (postα + postβ - 2.0)
}

// BinomPiPostMean returns Posterior mean of the Binomial proportion.
func BinomPiPostMean(α, β float64, n, k int64) float64 {
	var postα, postβ float64
	postα = α + float64(k)
	postβ = β + float64(n-k)
	return ((postα) / (postα + postβ))
}

/*
// BinomPiPostMedian returns Posterior median of the Binomial proportion.
func BinomPiPostMedian(α, β float64, n, k int64) float64 {
	// TO BE IMPLEMENTED
}
*/

// BinomPiPostVar returns Posterior variance of the Binomial proportion.
// Bolstad 2007 (2e): 151, eq. 8.5
func BinomPiPostVar(α, β float64, n, k int64) float64 {
	var postα, postβ float64
	postα = α + float64(k)
	postβ = β + float64(n-k)
	return (postα * postβ) / ((postα + postβ) * (postα + postβ) * (postα + postβ + 1.0))
}

// BinomPiPMS returns Posterior mean square of p (Binomial proportion).
// Bolstad 2007 (2e): 152-153, eq. 8.7
func BinomPiPMS(α, β float64, n, k, whichpi int64) float64 {
	const (
		mean = iota
		median
		modus
	)

	var postmean, postvar, pihat float64

	postvar = BinomPiPostVar(α, β, n, k)
	postmean = BinomPiPostMean(α, β, n, k)

	switch whichpi {
	case mean:
		pihat = BinomPiPostMean(α, β, n, k)
	case median:
		pihat = BinomPiPostMedian(α, β, n, k)
	case MODUS:
		pihat = BinomPiPostModus(α, β, n, k)
	}
	return postvar + (postmean-pihat)*(postmean-pihat)
}

// Binomial proportion, credible interval, beta prior, equal tail area.
// Bolstad 2007 (2e): 153
// untested ...
func BinomPiCrIBP(α, β, alpha float64, n, k int64) (low, upp float64) {
	// k-observed successes
	// n - total number of observations
	// α - beta prior a
	// β - beta prior b
	// alpha - posterior probability that the true proportion lies outside the credible interval

	low = dst.BetaQtlFor(alpha/2.0, α+float64(k), β+float64(n-k))
	upp = dst.BetaQtlFor(1.0-alpha/2.0, α+float64(k), β+float64(n-k))
	return
}

// BinomPiCrIBPriNApprox returns boundaries of the credible interval of theBinomial proportion, beta prior, equal tail area, normal approximation,
// Bolstad 2007 (2e): 154-155, eq. 8.8
// untested ...
func BinomPiCrIBPriNApprox(α, β, alpha float64, n, k int64) (low, upp float64) {
	// Arguments:
	// k - observed successes
	// n - total number of observations
	// a - beta prior a
	// b - beta prior b
	// alpha - posterior probability that the true proportion lies outside the credible interval
	//
	// Returns:
	// low, upp - lower and upper boundary of the credible interval

	var postmean, postvar, postα, postβ, z, low, upp float64

	postα = α + float64(k)
	postβ = β + float64(n-k)

	postmean = postα / (postα + postβ)
	postvar = (postα * postβ) / ((postα + postβ) * (postα + postβ) * (postα + postβ + 1.0))
	z = dst.ZQtlFor(alpha / 2)

	low = postmean - z*math.Sqrt(postvar)
	upp = postmean + z*math.Sqrt(postvar)
	return low, upp
}

// Binomial proportion, Likelihood
func BinomPiLike(pi float64, n, k int64) float64 {
	return math.Pow(pi, float64(k)) * math.Pow(1-pi, float64(n-k))
}

// BinomPiDeviance returns the Deviance of the Binomial proportion. 
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
	return dst.BetaNext(α+float64(k), β+float64(n-k))
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
