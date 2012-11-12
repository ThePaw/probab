// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Uniform (Flat) distribution. 
// The continuous uniform distribution or rectangular distribution is a family of probability distributions such that for each member of the family, all intervals of the same length on the distribution's support are equally probable. The support is defined by the two parameters, a and b, which are its minimum and maximum values. The distribution is often abbreviated U(a,b). It is the maximum entropy probability distribution for a random variate X under no constraint other than that it is contained in the distribution's support.
//
// Parameters: 
// a ∈ (-∞, b)		lower boundary (real)
// b ∈ (a, ∞)		upper boundary (real)
//
// Support: 
// x ∈ [a, b]		(real)

import (
	"math/rand"
	"math"
)

// UniformPDF returns the PDF of the Uniform distribution. 
func UniformPDF(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		if a <= x && x <= b {
			return 1/(b-a)
		}
		return 0
	}
}

// UniformLnPDF returns the natural logarithm of the PDF of the Uniform distribution. 
func UniformLnPDF(a, b float64)  func(x float64) float64 {
	return func(x float64) float64 {
		if a <= x && x <= a {
			return math.Log(1/(b-a))
		}
		return negInf
	}
}

// UniformPDFAt returns the value of PDF of Uniform distribution at x. 
func UniformPDFAt(a, b, x float64) float64 {
	pdf := UniformPDF(a, b)
	return pdf(x)
}

// UniformCDF returns the CDF of the Uniform distribution. 
func UniformCDF(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		switch {
		case x < a:
			return 0
		case x >b:
			return 1
		}
		return (x-a)/(b-a)
	}
}

// UniformCDFAt returns the value of CDF of the Uniform distribution, at x. 
func UniformCDFAt(a, b, x float64) float64 {
	cdf := UniformCDF(a, b)
	return cdf(x)
}

// UniformNext returns random number drawn from the Uniform distribution. 
func UniformNext(a, b float64)  float64 {  
	return a + (b-a)*rand.Float64()
}

// Uniform returns the random number generator with  Uniform distribution. 
func Uniform(a, b float64) func() float64 { 
	return func() float64 { return UniformNext(a, b) } 
}


// UniformMean returns the mean of the Uniform distribution. 
func UniformMean(a, b float64) float64 {
	return (a+b)/2
}


// UniformMedian returns the median of the Uniform distribution. 
func UniformMedian(a, b float64) float64 {
	return (a+b)/2
}

// UniformVar returns the variance of the Uniform distribution. 
func UniformVar(a, b float64) float64 {
	return (b-a)*(b-a)/12
}

// UniformStd returns the standard deviation of the Uniform distribution. 
func UniformStd(a, b float64) float64 {
	return (b-a)/3.4641016151377543
}

// UniformSkew returns the skewness of the Uniform distribution. 
func UniformSkew(a, b float64) (s float64) {
	return 0
}

// UniformExKurt returns the excess kurtosis of the Uniform distribution. 
func UniformExKurt(a, b float64) float64 {
	return -6.0/5
}

// UniformMGF returns the moment-generating function of the YYYYYYYYYYYYYYYYYYY distribution. 
func UniformMGF(a, b, t float64) float64 {
	return (math.Exp(t*b) - math.Exp(t*a)) / (t*(b-a))
}

