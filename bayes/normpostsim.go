// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Simulation from Bayesian normal sampling model.
// Ref.: Albert (2009)

/*
    a=prior$sigma2[1]
    b=prior$sigma2[2]
    mu0=prior$mu[1]
    tau2=prior$mu[2]
    S = sum((data - mean(data))^2)
    xbar = mean(data)
    n = length(data)

    SIGMA2=rep(0,m)
    MU=rep(0,m)
    sigma2=S/n
    for (j in 1:m)
    {
    prec=n/sigma2+1/tau2
    mu1=(xbar*n/sigma2+mu0/tau2)/prec
    v1=1/prec
    mu=rnorm(1,mu1,sqrt(v1))

    a1=a+n/2
    b1=b+sum((data-mu)^2)/2
    sigma2=rigamma(1,a1,b1)

    SIGMA2[j]=sigma2
    MU[j]=mu
    }
}
    return(list(mu = MU, sigma2 = SIGMA2))
}
*/

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
func NormPostSim(data []float64, a, b, mu0, tau2 float64, m int) (MU, SIGMA2 []float64) {
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
	//    S = sum((data - mean(data))^2)
	diff2 := make([]float64, n)
	for i, val := range data {
		diff2[i] = (val - xbar) * (val - xbar)
	}

	S := sum(diff2)
	SIGMA2 = make([]float64, m)
	MU = make([]float64, m)
	sigma2 := S / float64(n)

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

		SIGMA2[j] = sigma2
		MU[j] = mu
	}
	return
}

// NormPostSimNoPrior returns a simulated sample from the joint posterior distribution of the mean and variance for a normal
// sampling prior.
func NormPostSimNoPrior(data []float64, m int) (MU, SIGMA2 []float64) {
	xbar := mean(data)
	n := len(data)
	//    S = sum((data - mean(data))^2)
	diff2 := make([]float64, n)
	for i, val := range data {
		diff2[i] = (val - xbar) * (val - xbar)
	}

	S := sum(diff2)
	SIGMA2 = make([]float64, m)
	MU = make([]float64, m)
	for i, _ := range MU {
		SIGMA2[i] = S / dst.ChiSquareNext(int64(n)-1)
		sd := sqrt(SIGMA2[i]) / sqrt(float64(n))
		MU[i] = dst.NormalNext(xbar, sd)
	}
	return
}
