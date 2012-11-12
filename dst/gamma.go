// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Gamma distribution. 
// Parameters: 
// k > 0.0		shape parameter, 
// θ (Theta) > 0.0	scale parameter. 
// Alternatively, shape parameter α = k and an inverse scale parameter β = 1⁄θ, is called a rate parameter.
// If k is an integer, then the distribution represents an Erlang distribution; i.e., the sum of k  independent exponentially-distributed random variables, each of which has a mean of θ (which is equivalent to a rate parameter of 1/θ). Equivalently, if α is an integer, then the distribution again represents an Erlang distribution, i.e. the sum of α independent exponentially-distributed random variables, each of which has a mean of 1/β (which is equivalent to a rate parameter of β).
// Support: 
// x ∈ (0, ∞)

import (
	"fmt"
	"math"
	. "code.google.com/p/go-fn/fn"
)

// GammaPDF returns the PDF of the Gamma distribution. 
func GammaPDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return pow(x, k-1) * exp(-x/θ) / (Γ(k) * pow(θ, k))
	}
}

// GammaLnPDF returns the natural logarithm of the PDF of the Gamma distribution. 
func GammaLnPDF(α float64, λ float64) func(x float64) float64 {
	expPart := ExponentialLnPDF(λ)
	return func(x float64) float64 {
		if x < 0 {
			return negInf
		}
		return expPart(x) + (α-1)*log(λ*x) - LnΓ(α)
	}
}

// GammaPDFAt returns the value of PDF of Gamma distribution at x. 
func GammaPDFAt(k, θ, x float64)  float64 {
	pdf := GammaPDF(k , θ)
	return pdf(x)
}

// GammaCDF returns the CDF of the Gamma distribution. 
// Analytic solution, did not pass some tests!
func GammaCDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || θ < 0 {
			panic(fmt.Sprintf("k < 0 || θ < 0"))
		}
		if x < 0 {
			return 0
		}
		return Iγ(k, x/θ) / Γ(k)
	}
}

// GammaCDFint returns the CDF of the Gamma distribution, for integer k only. 
// Cumulative distribution function, for integer k only
func GammaCDFint(k int64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || θ < 0 {
			panic(fmt.Sprintf("k < 0 || θ < 0"))
		}
		if x < 0 {
			return 0
		}
		return Iγint(k, x/θ) / Γ(float64(k))
	}
}

/*
// Cumulative distribution function, using gamma incomplete integral  DOES NOT WORK !!!
func GammaCDF(k float64, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || θ < 0 {
			panic(fmt.Sprintf("k < 0 || θ < 0"))
		}
		if x < 0 {
			return 0
		}
		return IGam(θ, k*x)
	}
}
*/

// GammaCDFAt returns the value of CDF of the Gamma distribution, at x. 
func GammaCDFAt(k, θ, x float64)  float64 {
	cdf := GammaCDFint(int64(math.Ceil(k)) , θ)
	return cdf(x)
}

// GammaNext returns random number drawn from the Gamma distribution. 
func GammaNext(α float64, λ float64) float64 {
	//if α is a small integer, this way is faster on my laptop
	if α == float64(int64(α)) && α <= 15 {
		x := ExponentialNext(λ)
		for i := 1; i < int(α); i++ {
			x += ExponentialNext(λ)
		}
		return x
	}

	if α < 0.75 {
		return RejectionSample(GammaPDF(α, λ), ExponentialPDF(λ), Exponential(λ), 1)
	}

	//Tadikamalla ACM '73
	a := α - 1
	b := 0.5 + 0.5*sqrt(4*α-3)
	c := a * (1 + b) / b
	d := (b - 1) / (a * b)
	s := a / b
	p := 1.0 / (2 - exp(-s))
	var x, y float64
	for i := 1; ; i++ {
		u := UniformNext(0, 1)
		if u > p {
			var e float64
			for e = -log((1 - u) / (1 - p)); e > s; e = e - a/b {
			}
			x = a - b*e
			y = a - x
		} else {
			x = a - b*log(u/p)
			y = x - a
		}
		u2 := UniformNext(0, 1)
		if log(u2) <= a*log(d*x)-x+y/b+c {
			break
		}
	}
	return x / λ
}

// Gamma returns the random number generator with  Gamma distribution. 
func Gamma(α float64, λ float64) func() float64 {
	return func() float64 { return GammaNext(α, λ) }
}

// GammaQtl returns the inverse of the CDF (quantile) of the Gamma distribution. 
func GammaQtl(k, θ float64) func(x float64) float64 {
	return func(x float64) float64 {
		var eps, ynew, h float64
		eps = 1e-10
		y := k * θ
		yold := y
	L:
		for i := 0; i < 100; i++ {
			h = (GammaCDFAt(k, θ, yold) - x) / GammaPDFAt(k, θ, yold)
			ynew = yold - h
			if ynew <= eps {
				ynew = yold / 10
				h = yold - ynew
			}
			if math.Abs(h) < eps {
				break L
			}
			yold = ynew
		}
		return ynew
	}
}

// GammaQtlFor returns the inverse of the CDF (quantile) of the Gamma distribution, for given probability.
func GammaQtlFor(k, θ, p float64)  float64 {
	cdf:=GammaQtl(k, θ)
	return cdf(p)
}

// GammaMean returns the mean of the Gamma distribution. 
func GammaMean(k, θ float64) float64 {
	return k*θ
}

// GammaMode returns the mode of the Gamma distribution. 
func GammaMode(k, θ float64) float64 {
	if k <=1 {
		panic("mode not defined for k <=1")
	} 
	return (k-1)*θ
}

// GammaVar returns the variance of the Gamma distribution. 
func GammaVar(k, θ float64) float64 {
	return k*θ*θ
}

// GammaStd returns the standard deviation of the Gamma distribution. 
func GammaStd(k, θ float64) float64 {
	return math.Sqrt(k)*θ
}

// GammaSkew returns the skewness of the Gamma distribution. 
func GammaSkew(k, θ float64) float64 {
	return 2/math.Sqrt(k)
}

