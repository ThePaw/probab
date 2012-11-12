package stat

import . "code.google.com/p/probab/dst"

// One-sided (frequentist) Confidence Intervals for Observed "Nonconforming" Units in a Random Sample
// Source: Hahn, G. J., and W. Q. Meeker, "Statistical Intervals / A Guide for Practitioners," J. Wiley & Sons, New York.  1991.

func BinomPConfInt(n, k int64, alpha float64) (float64, float64) {

	/*
	Alpha	100(1-alpha) is the confidence 
	n	Sample size	
	k	Observed number of successes (p=n/k)
	lCL	Lower confidence limit
	uCL	Upper confidence limit
	*/


	var lCL, uCL float64
	if k <= 0 {
		lCL = 0.0
	} else {
		lCL = 1.0 / (1.0 + float64((n-k+1))*FQtlFor(2*n-2*k+2, 2*k, alpha)/float64(k))
	}

	if k >= n {
		uCL = 1.0
	} else {
		uCL = 1.0 / (1.0 + float64((n-k))/(float64((k+1))*FQtlFor(2*k+2, 2*n-2*k, alpha)))
	}
	return lCL, uCL
}

