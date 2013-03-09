// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Selection of Beta Prior Given Knowledge of Two Quantiles.
// Ref.: Albert (2009)

import (
	dst "code.google.com/p/probab/dst"
)

// betaprior1 returns the prior mean m given a beta(k*m, K*(1-m)) prior 	
// where the pth quantile is given by x. 
func betaprior1(k, x, p float64) float64 {
	var m0, p0 float64
	mLo := 0.0
	mHi := 1.0
	rep := true
	for rep == true {
		m0 = (mLo + mHi) / 2.0
		p0 = dst.BetaCDFAt(k*m0, k*(1-m0), x)
		if p0 < p {
			mHi = m0
		} else {
			mLo = m0
		}
		if abs(p0-p) < .0001 {
			rep = false
		}
	}
	return m0
}

//  BetaFromQtls finds the shape parameters of a beta density that matches knowledge of two quantiles of the distribution.
func BetaFromQtls(p1, x1, p2, x2 float64) (alpha, beta float64) {
	// p1	first probability
	// x1	its quantile
	// p2	second probability
	// x2	its quantile
	// alpha, beta	params of the corresponding Beta distribution.

	m := make([]float64, 100)
	logK := make([]float64, 100)
	k := make([]float64, 100)
	for i := 0; i < 100; i++ {
		v := -3.0 + 0.1111111*float64(i)
		logK[i] = v
		k[i] = exp(v)
		m[i] = betaprior1(exp(v), x1, p1)
	}

	prob2 := make([]float64, 100)
	for i, _ := range prob2 {
		prob2[i] = dst.BetaCDFAt(k[i]*m[i], k[i]*(1-m[i]), x2)
	}

	yOut := linInt(prob2, logK, p2)
	k0 := exp(yOut)
	m0 := betaprior1(k0, x1, p1)
	alpha = k0 * m0
	beta = k0 * (1 - m0)
	return
}
