// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Log posterior with Poisson sampling and gamma prior.
// Ref.: Albert (2009)

import (
	"code.google.com/p/probab/dst"
)

// LogPoissGamma returns the logarithm of the posterior density of a Poisson log mean with a gamma prior.
func LogPoissGamma(theta, y []float64, sh, rt float64) []float64 {
	// Arguments:
	// theta - vector of values of the log mean parameter
	// y - vector of observations, and 
	// sh, rt  -  shape and rate=1/scale parameters of the gamma prior

	// Returns:
	// logPosterior - vector of values of the log posterior for all values in theta

	lambda := make([]float64, len(theta))
	for i, _ := range lambda {
		lambda[i] = exp(theta[i])
	}

	// shape=sum(y)+1
	shape := 1.0
	for _, val := range y {
		shape += val
	}

	scale := 1 / float64(len(y))

	logLike := make([]float64, len(lambda))
	for i, _ := range logLike {
		logLike[i] = dst.GammaLnPDFAt(shape, scale, lambda[i])
	}

	logPrior := make([]float64, len(theta))
	for i, _ := range logPrior {
		logPrior[i] = dst.GammaLnPDFAt(sh, 1/(rt*lambda[i]), lambda[i])
	}

	logPosterior := make([]float64, len(logPrior))
	for i, _ := range logPrior {
		logPosterior[i] = logLike[i] + logPrior[i]
	}
	return logPosterior
}
