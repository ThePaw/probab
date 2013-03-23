// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Log posterior of Poisson/gamma exchangeable model.
// Ref.: Albert (2009)

func logG(y, e, alpha, beta float64) float64 {
	return lnΓ(alpha+y) - (y+alpha)*log(e+beta) + alpha*log(beta) - lnΓ(alpha)
}

// PoissGamExch returns the log posterior density of log alpha and log mu for a Poisson/gamma exchangeable model.
func PoissGamExch(theta1, theta2, z0 float64, y, e []float64) float64 {
	// Arguments:
	// theta1,  theta2 - parameter values of log alpha and log mu
	// z0, prior hyperparameter
	// data: y and e
	// Returns:
	// value of the log posterior

	alpha := exp(theta1)
	mu := exp(theta2)
	beta := alpha / mu
	val := 0.0
	for i, _ := range y {
		val += logG(y[i], e[i], alpha, beta)
	}
	val += log(alpha) - 2*log(alpha+z0)
	return val
}
