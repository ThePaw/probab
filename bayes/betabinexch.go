// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Log posterior of logit mean and log precision for Binomial/beta exchangeable model.
// Ref.: Albert (2009)

func logB(y, n, k, eta float64) float64 {
	return lnB(k*eta+y, k*(1-eta)+n-y) - lnB(k*eta, k*(1-eta))
}

// BetaBinExch returns the log posterior density of logit mean and log precision for a Binomial/beta exchangeable model.
func BetaBinExch(theta1, theta2 float64, y, n []float64) float64 {
	// Arguments:
	// theta1,  theta2 -  parameter values of logit eta and log k
	// data:
	//  y - counts
	// n - sample sizes
	// Returns:
	// value of the log posterior

	eta := exp(theta1) / (1 + exp(theta1))
	k := exp(theta2)
	val := 0.0
	for i, _ := range y {
		val += logB(y[i], n[i], k, eta)
	}
	val += theta2 - 2*log(1+exp(theta2))
	return val
}

// BetaBinExch0 returns the log posterior density of mean and precision for a Binomial/beta exchangeable model.
func BetaBinExch0(theta1, theta2 float64, y, n []float64) float64 {
	// Arguments:
	// theta1,  theta2 -  parameter values of logit eta and log k
	// data:
	//  y - counts
	// n - sample sizes
	// Returns:
	// value of the log posterior

	eta := theta1
	k := theta2
	val := 0.0
	for i, _ := range y {
		val += logB(y[i], n[i], k, eta)
	}
	val += -2*log(1+k) - log(eta) - log(1-eta)
	return val
}
