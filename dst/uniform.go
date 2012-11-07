// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Uniform (Flat) distribution

import (
	"math/rand"
	"math"
)

// Uniform_PDF returns the PDF of the Uniform distribution. 
func Uniform_PDF(a, b float64) func(x float64) float64 {
	return func(x float64) float64 {
		if a <= x && x <= b {
			return 1/(b-a)
		}
		return 0
	}
}

// Uniform_LnPDF returns the natural logarithm of the PDF of the Uniform distribution. 
func Uniform_LnPDF(a, b float64)  func(x float64) float64 {
	return func(x float64) float64 {
		if a <= x && x <= a {
			return math.Log(1/(b-a))
		}
		return negInf
	}
}

// Uniform_PDF_At returns the value of PDF of Uniform distribution at x. 
func Uniform_PDF_At(a, b, x float64) float64 {
	pdf := Uniform_PDF(a, b)
	return pdf(x)
}

// Uniform_CDF returns the CDF of the Uniform distribution. 
func Uniform_CDF(a, b float64) func(x float64) float64 {
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

// Uniform_CDF_At returns the value of CDF of the Uniform distribution, at x. 
func Uniform_CDF_At(a, b, x float64) float64 {
	cdf := Uniform_CDF(a, b)
	return cdf(x)
}

// NextUniform returns random number drawn from the Uniform distribution. 
func NextUniform(a, b float64)  float64 {  
	return a + (b-a)*rand.Float64()
}

// Uniform returns the random number generator with  Uniform distribution. 
func Uniform(a, b float64) func() float64 { 
	return func() float64 { return NextUniform(a, b) } 
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

