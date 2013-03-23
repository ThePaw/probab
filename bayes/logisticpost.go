// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Log posterior for a binary response model with a logistic link and a uniform prior.
// Ref.: Albert (2009)

func logFn(x, n, y, beta0, beta1 float64) float64 {
	lp := beta0 + beta1*x
	p := exp(lp) / (1 + exp(lp))
	return y*log(p) + (n-y)*log(1-p)
}

// LogisticPost returns the log posterior density of (beta0, beta1) when yi are independent binomial(ni, pi)
// and logit(pi)=beta0+beta1*xi and a uniform prior is placed on (beta0, beta1).
func LogisticPost(x, n, y []float64, beta0, beta1 float64) float64 {

	val := 0.0
	for i, _ := range y {
		val += logFn(x[i], n[i], y[i], beta0, beta1)
	}
	return val
}
