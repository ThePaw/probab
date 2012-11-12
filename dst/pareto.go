// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Pareto Type I distribution (sometimes referred to as the Bradford distribution). 
//
// Parameters: 
// θ > 0.0 (scale) 
// α > 0.0 (shape) 
//
// Support: 
// k x >= θ 
// x ∈ (0, ∞)

import (
	"math"
	. "code.google.com/p/go-fn/fn"
)

// ParetoChkParams checks parameters of the Pareto Type I distribution. 
func ParetoChkParams(θ, α float64) bool {
	ok := true
	if α <= 0 || θ <= 0  {
		ok = false
	}
	return ok
}

// ParetoChkSupport checks support of the Pareto Type I distribution. 
func ParetoChkSupport(x float64) bool {
	ok := true
	if x < 0 {
		ok = false
	}
	return ok
}

// ParetoPDF returns the PDF of the Pareto Type I distribution. 
func ParetoPDF(θ, α float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < θ {
			return 0
		}
		t1 := α*math.Pow(θ, α)
		t2 := math.Pow(x, α + 1)
		return t1/t2
	}
}

// ParetoPDFAt returns the value of PDF of Pareto Type I distribution at x. 
func ParetoPDFAt(θ, α, x float64) float64 {
	pdf := ParetoPDF(θ, α)
	return pdf(x)
}

// ParetoCDF returns the CDF of the Pareto Type I distribution. 
func ParetoCDF(θ, α float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < θ {
			return 0
		}
		return 1-math.Pow(θ/x, α)
	}
}

// ParetoCDFAt returns the value of CDF of the Pareto Type I distribution, at x. 
func ParetoCDFAt(θ, α, x float64) float64 {
	cdf := ParetoCDF(θ, α)
	return cdf(x)
}

// ParetoQtl returns the inverse of the CDF (quantile) of the Pareto Type I distribution. 
func ParetoQtl(θ, α float64) func(p float64) float64 {
	return func(p float64) float64 {
		return math.Pow(θ*(1-p),(-1/α))
	}
}

// ParetoQtlFor returns the inverse of the CDF (quantile) of the Pareto Type I distribution, for given probability.
func ParetoQtlFor(θ, α, p float64) float64 {
	cdf := ParetoQtl(θ, α)
	return cdf(p)
}

// ParetoMean returns the mean of the Pareto Type I distribution. 
func ParetoMean(θ, α float64) float64 {
	if α <= 1 {
		return posInf
	}
	return α*θ/(α-1)
}

// ParetoMode returns the mode of the Pareto Type I distribution. 
func ParetoMode(θ, α float64) float64 {
	return θ
}

// ParetoMedian returns the median of the Pareto Type I distribution. 
func ParetoMedian(θ, α float64) float64 {
	return θ * math.Pow(2, 1/α)
}

// ParetoVar returns the variance of the Pareto Type I distribution. 
func ParetoVar(θ, α float64) float64 {
	if α <= 2 {
		return posInf
	}
	return θ*θ*α/((α-1)*(α-1)*(α-2))
}

// ParetoStd returns the standard deviation of the Pareto Type I distribution. 
func ParetoStd(θ, α float64) float64 {
	if α <= 2 {
		return posInf
	}
	return θ/(α-1)*math.Sqrt(α/(α-2))
}

// ParetoSkew returns the skewness of the Pareto Type I distribution. 
func ParetoSkew(θ, α float64) float64 {
	if α <= 3 {
		panic("skewness not defined for α <= 3")
	}
	return 2*(1+α)/(α-3)*math.Sqrt((α-2)/α)
}

// ParetoExKurt returns the excess kurtosis of the Pareto Type I distribution. 
func ParetoExKurt(θ, α float64) float64 {
	if α <= 4 {
		panic("skewness not defined for α <= 4")
	}
	return 6*(α*α*α+α*α-6*α-2) /(α*(α-3)*(α-4))
}

// ParetoMGF returns the moment-generating function of the Pareto Type I distribution. 
func ParetoMGF(θ, α, t float64) float64 {
	if t >= 0 {
		panic("MGF not defined for t >= 0")
	}
	return α*math.Pow((-θ*t), α) * IΓ(-α, -θ*t)
}
