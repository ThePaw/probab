// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Posterior for normal sampling and a mixture of normals prior

import (
	"code.google.com/p/probab/dst"
)

// NormalNormalMix returns the parameters and mixing probabilities for a normal sampling problem, variance known,
// where the prior is a discrete mixture of normal densities.
func NormalNormalMix(probs, priorMean, priorVar []float64, y, sigma2 float64) (postProbs, postMean, postVar []float64) {
	// Arguments:
	// probs vector of probabilities of the normal components of the prior
	// priorMean vector of the prior means for a normal component of the prior
	// priorVar vector of the prior variances for a normal component of the prior
	// y -  observation
	// sigma2 - sampling variance
	// Returns:
	// postProbs - vector of probabilities of the normal components of the posterior
	// postMean  - vector of means  for a normal component of the posterior
	// postMean  - vector of variances  for a normal component of the posterior

	postPrecision := make([]float64, len(probs))
	for i, _ := range probs {
		postPrecision[i] = 1/priorVar[i] + 1/sigma2
	}

	postVar = make([]float64, len(probs))
	for i, _ := range probs {
		postVar[i] = 1 / postPrecision[i]
	}

	postMean = make([]float64, len(probs))
	for i, _ := range probs {
		postMean[i] = (y/sigma2 + priorMean[i]/priorVar[i]) / postPrecision[i]
	}

	mProb := make([]float64, len(probs))
	for i, _ := range probs {
		mProb[i] = dst.NormalPDFAt(priorMean[i], sqrt(sigma2+priorVar[i]), y)
	}

	prod := make([]float64, len(probs))
	for i, _ := range probs {
		prod[i] = probs[i] * mProb[i]
	}

	postProbs = make([]float64, len(probs))
	for i, _ := range probs {
		postProbs[i] = probs[i] * mProb[i] / sum(prod)
	}

	return postProbs, postMean, postVar
}
