// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

// D’Agostino test for skewness in normally distributed data.

package stat

import (
	"code.google.com/p/probab/dst"
	"sort"
)

// Agostino performs D’Agostino test for skewness in normally distributed data vector.
func Agostino(x []float64, alternative int) (s3, z, pval float64) {
	// Arguments: 
	// x - vector of observations
	// alternative - 0 = "twoSided", 1 = "less", 2 = "greater"
	//
	// Details: 
	// Under the hypothesis of normality, data should be symmetrical (i.e. skewness should be equal to
	// zero). This test has such null hypothesis and is useful to detect a significant skewness in normally
	// distributed data.
	//
	// Returns: 
	// A list with class htest containing the following components:
	// s3 - skewness estimator 
	// z - its transformation
	// pval - the p-value for the test.

	const (
		twoSided = iota
		less
		greater
	)

	sort.Float64s(x)
	n := float64(len(x))
	dm := diffMean(x)
	d3 := make([]float64, len(dm))
	for i, val := range dm {
		d3[i] = val * val * val
	}

	d2 := make([]float64, len(dm))
	for i, val := range dm {
		d2[i] = val * val
	}

	//s3 <- (sum((x-mean(x))^3)/n)/(sum((x-mean(x))^2)/n)^(3/2)

	s3 = (sum(d3) / n) / pow((sum(d2)/n), 1.5)
	y := s3 * sqrt((n+1)*(n+3)/(6*(n-2)))
	b2 := 3 * (n*n + 27*n - 70) * (n + 1) * (n + 3) / ((n - 2) * (n + 5) * (n + 7) * (n + 9))
	w := sqrt(-1 + sqrt(2*(b2-1)))
	d := 1 / sqrt(log10(w))
	a := sqrt(2 / (w*w - 1))
	z = d * log10(y/a+sqrt((y/a)*(y/a)+1))
	pval = 1 - dst.NormalCDFAt(0, 1, z)

	switch alternative {
	case twoSided:
		pval = 2 * pval
		if pval > 1 {
			pval = 2 - pval
		}
	case less: // do nothing
	case greater:
		pval = 1 - pval
	}
	return s3, z, pval
}
