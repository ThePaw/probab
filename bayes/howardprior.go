package bayes

import (
	"math/rand"
)

// LnHowardPrior returns the logarithm of a dependent prior on two proportions proposed by 
// Howard in a Statistical Science paper in 1998.
func LnHowardPrior(p1, p2, alpha, beta, gamma, delta, sigma float64) float64 {
	// Arguments:
	// p1 , p2 - proportions to be tested
	// par vector containing parameter values alpha, beta, gamma, delta, sigma
	// Returns: 
	// z - logarithm of a dependent prior

	u := log(p1/(1-p1)*(1-p2)/p2) / sigma
	z := -0.5*u*u + (alpha-1)*log(p1) + (beta-1)*log(1-p1) + (gamma-1)*log(p2) + (delta-1)*log(1-p2)
	return z
}

// HowardPosteriorProb returns the posterior probability that p1 > p2.
func HowardPosteriorProb(y1, n1, y2, n2, alpha, beta, gamma, delta, sigma float64) float64 {
	nIter := 10000000
	// updated params for posterior that has the same functional form as prior
	alpha += y1
	beta += n1 - y1
	gamma += y2
	delta += n2 - y2
	sum1 := 0.0
	sum2 := 0.0
	// brute force sampling from posterior, should be improved
	for i := 0; i < nIter; i++ {
		p1 := rand.Float64()
		if p1 < 0.0001 {
			p1 = 0.0001
		}
		if p1 > 0.9999 {
			p1 = 0.9999
		}

		p2 := rand.Float64()
		if p2 < 0.0001 {
			p2 = 0.0001
		}
		if p2 > 0.9999 {
			p2 = 0.9999
		}
		// posterior
		pr := exp(LnHowardPrior(p1, p2, alpha, beta, gamma, delta, sigma))
		if p1 > p2 {
			sum1 += pr
		} else {
			sum2 += pr
		}
	}
	return sum1 / (sum1 + sum2)
}
