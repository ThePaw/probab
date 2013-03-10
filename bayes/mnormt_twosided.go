// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

//Bayesian test of a two-sided hypothesis about a normal mean
// Ref.: Albert (2009): 184 [mnormt.onesided()]

// NormMeanTestTwoSided does a Bayesian test that a normal mean is equal to a specified value using a normal prior.
func NormMeanTestTwoSided(m0, prob float64, t []float64, smpMean float64, smpSize int, popSd float64) (bf, post []float64) {
	// 
	// Arguments:
	// m0 - value of the normal mean to be tested
	// prob - prior probability of the hypothesis
	// t -  vector of values of the prior standard deviation under the alternative hypothesis
	// smpMean - sample mean
	// n - sample size
	// popSd - known value of the population standard deviation
	// 
	// Returns:
	// bf vector of values of the Bayes factor in support of the null hypothesis
	// post vector of posterior probabilities of the null hypothesis
	// 
	n := float64(smpSize)
	popVar := popSd * popSd
	diffMean2 := (smpMean - m0) * (smpMean - m0)
	num := 0.5*log(n) - log(popSd) - 0.5*n/popVar*diffMean2
	for i, val := range t {
		t[i] *= val
	}
	den := make([]float64, len(t))
	for i, _ := range den {
		den[i] = -0.5*log(popVar/n+t[i]) - 0.5/(popVar/n+t[i])*diffMean2
	}
	bf = make([]float64, len(t))
	for i, _ := range bf {
		bf[i] = exp(num - den[i])
	}
	post = make([]float64, len(t))
	for i, _ := range post {
		post[i] = prob * bf[i] / (prob*bf[i] + 1 - prob)
	}
	return
}
