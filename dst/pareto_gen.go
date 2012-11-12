// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Generalized Pareto distribution. 
// Klugman, S. A., Panjer, H. H. and Willmot, G. E. (2008), Loss Models, From Data to Decisions, Third Edition, Wiley.
//
// Parameters: 
// shape1 > 0.0
// shape2 > 0.0
// scale > 0.0
//
// Support: 
// x ... to be completed

import (
	"math"
	"math/rand"
	. "code.google.com/p/go-fn/fn"
)

// ParetoGPDF returns the PDF of the Generalized Pareto distribution. 
func ParetoGPDF(shape1, shape2, scale float64) func(x float64) (p float64) {
	// We work with the density expressed as
	// u^shape2 * (1 - u)^shape1 / (x * B(shape1, shape2))
	// with u = v/(1 + v) = 1/(1 + 1/v), v = x/scale.
	return func(x float64) (p float64) {
		if x < 0 {
			p = 0
		} else if x == 0 {
			if shape2 < 1 {
				p = math.Inf(1)
			} else if shape2 > 1 {
				p = 0
			} else {
				p = 1 / (scale * B(shape2, shape1))
			}
		}
		tmp := math.Log(x) - math.Log(scale)
		logu := - math.Log1p(math.Exp(-tmp))
		log1mu := - math.Log1p(math.Exp(tmp))
		p =  math.Exp(shape2 * logu + shape1 * log1mu - math.Log(x) - LnB(shape2, shape1))
		return
	}
}

// ParetoGPDFAt returns the value of PDF of Generalized Pareto distribution at x. 
func ParetoGPDFAt(shape1, shape2, scale, x float64) float64 {
	pdf := ParetoGPDF(shape1, shape2, scale)
	return pdf(x)
}

// ParetoGCDF returns the CDF of the Generalized Pareto distribution. 
func ParetoGCDF(shape1, shape2, scale float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		u := math.Exp(-math.Log1p(math.Exp(math.Log(scale) - math.Log(x))))
		cdf := BetaCDF(shape2, shape1)
		return cdf(u)
	}
}

// ParetoGCDFAt returns the value of CDF of the Generalized Pareto distribution, at x. 
func ParetoGCDFAt(shape1, shape2, scale, x float64) float64 {
	cdf := ParetoGCDF(shape1, shape2, scale)
	return cdf(x)
}

// ParetoGQtl returns the inverse of the CDF (quantile) of the Generalized Pareto distribution. 
func ParetoGQtl(shape1, shape2, scale float64) func(p float64) float64 {
	return func(p float64) float64 {
		if p < 0 || p > 1 {
			panic("bad param")
		}
		qtl := BetaQtl(shape2, shape1)
		return scale / (1.0 / qtl(p) - 1.0)
	}
}

// ParetoGQtlFor returns the inverse of the CDF (quantile) of the Generalized Pareto distribution, for given probability.
func ParetoGNext(shape1, shape2, scale float64) float64 {
	qtl := ParetoGQtl(shape1, shape2, scale)
	p := rand.Float64()
	return qtl(p)
}

// ParetoGMoment returns the n-th moment of the Generalized Pareto distribution. 
func ParetoGMoment(shape1, shape2, scale float64, order int) (x float64) {
	o := float64(order)
	if o  <= -shape2 || o  >= shape1 {
		x = math.Inf(1)
	} else {
		x = math.Pow(scale, o) * B(shape1 - o, shape2 + o) / B(shape1, shape2)
	}
	return
}

// ParetoGMean returns the mean of the Generalized Pareto distribution. 
func ParetoGMean(shape1, shape2, scale float64) float64 {
	return ParetoGMoment(shape1, shape2, scale, 1)
}

// ParetoGVar returns the variance of the Generalized Pareto distribution. 
func ParetoGVar(shape1, shape2, scale float64) float64 {
	return ParetoGMoment(shape1, shape2, scale, 2)
}

// ParetoGSkew returns the skewness of the Generalized Pareto distribution. 
func ParetoGSkew(shape1, shape2, scale float64) float64 {
	return ParetoGMoment(shape1, shape2, scale, 3)
}

// ParetoGExKurt returns the excess kurtosis of the Generalized Pareto distribution. 
func ParetoGExKurt(shape1, shape2, scale float64) float64 {
	return ParetoGMoment(shape1, shape2, scale, 4)
}

