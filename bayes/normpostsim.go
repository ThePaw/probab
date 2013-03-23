// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Simulation from Bayesian normal sampling model.
// Ref.: Albert (2009)

import (
	"code.google.com/p/probab/dst"
)

func rigamma(shape, rate float64) float64 {
	return (1 / dst.GammaNext(shape, 1/rate))
}

// NormPostSim returns a simulated sample from the joint posterior distribution of the mean and variance for a normal
// sampling prior with a noninformative or informative prior. The prior assumes mu and sigma2 are
// independent with mu assigned a normal prior with mean mu0 and variance tau2, and sigma2 is
// assigned a inverse gamma prior with parameters a and b.
func NormPostSim(data []float64, a, b, mu0, tau2 float64, m int) (postMu, postS2 []float64) {
	// Arguments:
	// data - vector of observations
	// prior params:
	//     a
	//     b
	//     mu0
	//     tau2
	// m - number of simulations desired
	// Returns:
	// mu - vector of simulated draws of normal mean
	// sigma2 - vector of simulated draws of normal variance

	xbar := mean(data)
	n := len(data)
	diff2 := make([]float64, n)
	for i, val := range data {
		diff2[i] = (val - xbar) * (val - xbar)
	}

	s := sum(diff2)
	postS2 = make([]float64, m)
	postMu = make([]float64, m)
	sigma2 := s / float64(n)

	for j := 0; j < m; j++ {
		prec := float64(n)/sigma2 + 1/tau2
		mu1 := (xbar*float64(n)/sigma2 + mu0/tau2) / prec
		v1 := 1 / prec
		//    mu=rnorm(1,mu1,sqrt(v1))
		mu := dst.NormalNext(mu1, sqrt(v1))

		a1 := a + float64(n)/2

		d2 := make([]float64, n)
		for i, val := range data {
			d2[i] = (val - mu) * (val - mu)
		}

		b1 := b + sum(d2)/2
		sigma2 := rigamma(a1, b1)

		postS2[j] = sigma2
		postMu[j] = mu
	}
	return
}

// NormPostSimNoPrior returns a simulated sample from the joint posterior distribution of the mean and variance for a normal
// sampling prior.
func NormPostSimNoPrior(data []float64, m int) (postMu, postS2 []float64) {
	// Arguments:
	// data - vector of observations
	// m - number of simulations desired
	// Returns:
	// mu - vector of simulated draws of normal mean
	// sigma2 - vector of simulated draws of normal variance

	xbar := mean(data)
	n := len(data)
	diff2 := make([]float64, n)

	for i, val := range data {
		diff2[i] = (val - xbar) * (val - xbar)
	}

	s := sum(diff2)
	postS2 = make([]float64, m)
	postMu = make([]float64, m)

	for i, _ := range postMu {
		postS2[i] = s / dst.ChiSquareNext(int64(n)-1)
		sd := sqrt(postS2[i]) / sqrt(float64(n))
		postMu[i] = dst.NormalNext(xbar, sd)
	}
	return
}

// NormPostNoPriorNext returns a  sampled tuple from the joint posterior distribution of the mean and variance for a normal
// sampling prior.
func NormPostNoPriorNext(data []float64) (postMu, postS2 float64) {
	// Arguments:
	// data - vector of observations
	// Returns:
	// postMu -  simulated draw of normal mean
	// postS2 -  simulated draw of normal variance

	xbar := mean(data)
	n := len(data)
	diff2 := make([]float64, n)

	for i, val := range data {
		diff2[i] = (val - xbar) * (val - xbar)
	}

	s := sum(diff2)
	postS2 = s / dst.ChiSquareNext(int64(n)-1)
	sd := sqrt(postS2) / sqrt(float64(n))
	postMu = dst.NormalNext(xbar, sd)
	return
}

// NormPostInfPriorNext returns a simulated tuple from the joint posterior distribution of the mean and variance for a normal
// sampling prior with a noninformative or informative prior. The prior assumes mu and sigma2 are
// independent with mu assigned a normal prior with mean mu0 and variance tau2, and sigma2 is
// assigned a inverse gamma prior with parameters a and b.
func NormPostInfPriorNext(data []float64, a, b, mu0, tau2 float64) (postMu, postS2 float64) {
	// Arguments:
	// data - vector of observations
	// prior params:
	//     a
	//     b
	//     mu0
	//     tau2
	// Returns:
	// postMu -  simulated draw of normal mean
	// postS2 -  simulated draw of normal variance

	xbar := mean(data)
	n := len(data)
	diff2 := make([]float64, n)
	for i, val := range data {
		diff2[i] = (val - xbar) * (val - xbar)
	}

	s := sum(diff2)
	postS2 = s / float64(n)
	prec := float64(n)/postS2 + 1/tau2
	mu1 := (xbar*float64(n)/postS2 + mu0/tau2) / prec
	v1 := 1 / prec
	postMu = dst.NormalNext(mu1, sqrt(v1))

	a1 := a + float64(n)/2

	d2 := make([]float64, n)
	for i, val := range data {
		d2[i] = (val - postMu) * (val - postMu)
	}

	b1 := b + sum(d2)/2
	postS2 = rigamma(a1, b1)
	return
}
