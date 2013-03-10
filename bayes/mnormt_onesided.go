// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"code.google.com/p/probab/dst"
)

//Bayesian test of one-sided hypothesis about a normal mean
// Ref.: Albert (2009): 184 [mnormt.onesided()]

// NormMeanTestOneSided does a Bayesian test of the hypothesis that a normal mean is less than or equal to a specified value.
func NormMeanTestOneSided(m0, priMean, priSD, smpMean float64, smpSize int, popSd float64) (bf, priOdds, postOdds, postH float64) {
	// 
	// Arguments
	// m0 - value of the normal mean to be tested
	// priMean - mean of the normal prior distribution
	// priSD -  standard deviation of the normal prior distribution
	// smpMean - sample mean
	// smpSize - sample size
	// popSd - known value of the population standard deviation
	// 
	// Returns
	// bf Bayes factor in support of the null hypothesis
	// priOdds prior odds of the null hypothesis
	// postOdds posterior odds of the null hypothesis
	// postH posterior probability of the null hypothesis
	// 
	// 
	n := float64(smpSize)
	priVar := priSD * priSD
	priH := dst.NormalCDFAt(priMean, priSD, m0)
	priA := 1 - priH
	priOdds = priH / priA

	popVar := popSd * popSd
	postPrecision := 1/priVar + n/popVar
	postVar := 1 / postPrecision
	postSd := sqrt(postVar)
	postMean := (smpMean*n/popVar + priMean/priVar) / postPrecision
	postH = dst.NormalCDFAt(postMean, postSd, m0)
	postA := 1 - postH
	postOdds = postH / postA
	bf = postOdds / priOdds
	return
}
