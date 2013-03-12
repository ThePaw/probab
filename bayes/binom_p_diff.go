// Bayesian inference for the difference between two binomial proportions using the Normal (Gaussian) approximation.
// Bolstad 2007 (2e): Chapter 13.4: 248-249.

// Compare the proportions of certain attribute in two populations. The true proportions are pi1 and pi2, unknown.
// We take a random sample from each of the populations and observe y1, y2 ... number of instances having the attribute.
// The distribution y1|pi1 is binomial(n1, pi1), similarly for y2|pi2, and they are independent.
// Let there be independent priors for pi1 ... beta(a1, b1), and similarly for pi2.
// Posterior for pi1 is beta(a1post, b1post), where a1post = a1 +y1, and b1post = b1 + n1 -y1, similarly for pi2.
// Approximate each posterior distribution with normal distribution having the same mean and variance as the beta.
// The posterior of pid = pi1 - pi2 is approximately normal(mdpost, vardpost), where:
// mdpost = a1post/(a1post+b1post) - a2post/(a2post+b2post), and
// vardpost = a1post*b1post/math.Sqrt(a1post+b1post)*(a1post+b1post+1)  +  a2post*b2post/math.Sqrt(a2post+b2post)*(a2post+b2post+1) 

package bayes

import (
	. "code.google.com/p/probab/dst"
	"math"
)

// Mean of posterior distribution of unknown difference of binomial proportions, approximated by Normal distribution
// Bolstad 2007 (2e): 248.
// untested ...
func BinomPiDiffMeanNApprox(a1, b1, a2, b2 float64, n1, n2, y1, y2 int64) float64 {

	a1post := a1 + float64(y1)
	b1post := b1 + float64(n1-y1)
	a2post := a2 + float64(y2)
	b2post := b2 + float64(n2-y2)

	return a1post/(a1post+b1post) - a2post/(a2post+b2post)
}

// Variance of posterior distribution of unknown difference of binomial proportions, approximated by Normal distribution
// Bolstad 2007 (2e): 248.
// untested ...
func BinomPiDiffVarNApprox(a1, b1, a2, b2 float64, n1, n2, y1, y2 int64) float64 {

	a1post := a1 + float64(y1)
	b1post := b1 + float64(n1-y1)
	a2post := a2 + float64(y2)
	b2post := b2 + float64(n2-y2)

	return a1post*b1post/math.Sqrt(a1post+b1post)*(a1post+b1post+1) + a2post*b2post/math.Sqrt(a2post+b2post)*(a2post+b2post+1)
}

// Credible interval for difference between binomial proportions, approximated by Normal distribution
// Bolstad 2007 (2e): 248, eq. 13.13
// postdiffmu = binomdiffpropnormapproxmu()
// postdiffsigma = sqrt(binomdiffpropnormapproxvar())
// untested ...
func BinomPiDiffCrI(postdiffmu, postdiffsigma, alpha float64) (float64, float64) {
	// postdiffmu		posterior mean for difference of normal means
	// postdiffsigma	posterior standard deviation for difference of normal means
	// alpha			posterior probability that the true mean lies outside the credible interval

	z := ZQtlFor(alpha / 2)
	low := postdiffmu - z*postdiffsigma
	high := postdiffmu + z*postdiffsigma
	return low, high
}

// One-sided test for difference between binomial proportions, approximated by Normal distribution
// Bolstad 2007 (2e): 248-249, eq. 13.14
// H0: mud <= 0 vs H1: mud > 0
// Note: The alternative is in the direction we wish to detect, and is what we want to detect.

func BinomPiDiffOneSidedP(postdiffmu, postdiffsigma float64) float64 {
	return ZCDFAt(-postdiffmu / postdiffsigma)
}

/*
Two-sided test for difference between binomial proportions, approximated by Normal distribution  ///// check it vs the book!!!
Bolstad 2007 (2e): 249
H0: mu1 - mu2 == 0 vs H1: mu1 - mu2 != 0
func BinomDiffPropTwoSidedProb(postdiffmu, postdiffsigma, alpha){
	low, high = normdiffmeansknowCrI(postdiffmu, postdiffsigma, alpha)

	if 0 < low || 0 > high return(REJECT) else return(ACCEPT)
}
*/
