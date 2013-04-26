// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Probability of a one sided null hypothesis from a sample from a posterior density.

// PNullSmpLowT returns the lower tail probability of a one sided null hypothesis from a sample from a posterior density.
func PNullSmpLowT(θ []float64, θ0 float64) float64 {
	return eCDF(θ, θ0)
}

// PNullSmpUppT returns the upper tail probability of a one sided null hypothesis from a sample from a posterior density.
func PNullSmpUppT(θ []float64, θ0 float64) float64 {
	return 1 - eCDF(θ, θ0)
}
