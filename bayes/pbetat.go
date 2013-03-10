// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// ProbBetaTest does a Bayesian test that a proportion is equal to a specified value using a beta prior.
func ProbBetaTest(p0, prob, a, b float64, succ, fail int) (bf, post float64) {
	// Arguments
	// p0 value of the proportion to be tested
	// prob prior probability of the hypothesis
	// a alpha parameter of the beta prior under the alternative hypothesis
	// b beta parameter of the beta prior under the alternative hypothesis
	// succ  number of successes
	// fail  number of failures
	// Value
	// bf the Bayes factor in support of the null hypothesis
	// post the posterior probability of the null hypothesis
	// Ref.: Albert (2009): 54.

	s := float64(succ)
	f := float64(fail)
	lbf := s*log(p0) + f*log(1-p0) + lnB(a, b) - lnB(a+s, b+f)
	bf = exp(lbf)
	post = prob * bf / (prob*bf + 1 - prob)
	return
}
