// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Gamma distribution. 
// Parameters: 
// k > 0.0		shape parameter, 
// λ (Theta) > 0.0	scale parameter. 
// Alternatively, shape parameter α = k and an inverse scale parameter β = 1⁄λ, is called a rate parameter.
// If k is an integer, then the distribution represents an Erlang distribution; i.e., the sum of k  independent exponentially-distributed random variables, each of which has a mean of λ (which is equivalent to a rate parameter of 1/λ). Equivalently, if α is an integer, then the distribution again represents an Erlang distribution, i.e. the sum of α independent exponentially-distributed random variables, each of which has a mean of 1/β (which is equivalent to a rate parameter of β).
// Support: 
// x ∈ (0, ∞)

import (
	. "code.google.com/p/go-fn/fn"
	"math"
)

// GammaPDF returns the PDF of the Gamma distribution. 
func GammaPDF(k float64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		return pow(x, k-1) * exp(-x/λ) / (Γ(k) * pow(λ, k))
	}
}

// GammaPDFAt returns the value of PDF of Gamma distribution at x. 
func GammaPDFAt(k, λ, x float64) float64 {
	pdf := GammaPDF(k, λ)
	return pdf(x)
}

// GammaCDF returns the CDF of the Gamma distribution. // TO BE REIMPLEMENTED
func GammaCDF(α float64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if isNaN(x) || isNaN(α) || isNaN(λ) {
			return x + α + λ
		}
		if α < 0. || λ <= 0. {
			return math.NaN()
		}
		if x == 0 {
			return 0
		}
		x /= λ
		if isNaN(x) { /* eg. original x = λ = +Inf */
			return x
		}
		if α == 0. { /* limit case; useful e.g. in pnchisq() */
			//	return (x <= 0) ? R_DT_0: R_DT_1; /* <= assert  pgamma(0,0) ==> 0 */
			if x <= 0 {
				return 0
			} else {
				return 1
			}
		}
		return pgamma_raw(x, α)
	}
}

// GammaCDFAt returns the value of CDF of the Gamma distribution, at x. 
func GammaCDFAt(k, λ, x float64) float64 {
	cdf := GammaCDF(k, λ)
	return cdf(x)
}

func GammaLnPDF(k float64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if x < 0 {
			return negInf
		}
		//		return log(pow(x, k-1) * exp(-x/λ) / (Γ(k) * pow(λ, k)))
		//		return log(pow(x, k-1)) + log(exp(-x/λ)) - log(Γ(k)) - log(pow(λ, k))
		//		return log(pow(x, k-1)) -x/λ - log(Γ(k)) - log(pow(λ, k))
		//		return (k-1)*log(x) -x/λ - log(Γ(k)) - k* log(λ)
		return (k-1)*log(x) - x/λ - LnΓ(k) - k*log(λ)
	}
}

// GammaLnPDFAt returns the value of PDF of Gamma distribution at x. 
func GammaLnPDFAt(α, λ, x float64) float64 {
	pdf := GammaLnPDF(α, λ)
	return pdf(x)
}

// GammaLnCDF returns the value of CDF of the Gamma distribution, at x. 
func GammaLnCDF(α float64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if isNaN(x) || isNaN(α) || isNaN(λ) {
			return x + α + λ
		}
		if α < 0. || λ <= 0. {
			return NaN
		}
		if x == 0 {
			return negInf
		}
		x /= λ
		if isNaN(x) { /* eg. original x = λ = +Inf */
			return x
		}
		if α == 0 { /* limit case; useful e.g. in pnchisq() */
			//	return (x <= 0) ? R_DT_0: R_DT_1; /* <= assert  pgamma(0,0) ==> 0 */
			if x <= 0 {
				return negInf
			} else {
				return 0
			}
		}
		return pgamma_raw_ln(x, α)
	}
}

// GammaLnCDFAt returns the value of CDF of the Gamma distribution, at x. 
func GammaLnCDFAt(k, λ, x float64) float64 {
	cdf := GammaLnCDF(k, λ) // TO BE IMPLEMENTED
	return cdf(x)
}

// GammaQtl returns the inverse of the CDF (quantile) of the Gamma distribution. 
func GammaQtl(α, λ float64) func(p float64) float64 {
	return func(p float64) float64 {
		var eps, ynew, h float64
		if p == 0 {
			return 0
		}
		if p == 1 {
			return posInf
		}

		eps = 1e-10
		y := α * λ
		yold := y
	L:
		for i := 0; i < 100; i++ {
			h = (GammaCDFAt(α, λ, yold) - p) / GammaPDFAt(α, λ, yold)
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
func GammaQtlFor(k, λ, p float64) float64 {
	cdf := GammaQtl(k, λ)
	return cdf(p)
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

// GammaMean returns the mean of the Gamma distribution. 
func GammaMean(k, λ float64) float64 {
	return k * λ
}

// GammaMode returns the mode of the Gamma distribution. 
func GammaMode(k, λ float64) float64 {
	if k <= 1 {
		panic("mode not defined for k <=1")
	}
	return (k - 1) * λ
}

// GammaVar returns the variance of the Gamma distribution. 
func GammaVar(k, λ float64) float64 {
	return k * λ * λ
}

// GammaStd returns the standard deviation of the Gamma distribution. 
func GammaStd(k, λ float64) float64 {
	return math.Sqrt(k) * λ
}

// GammaSkew returns the skewness of the Gamma distribution. 
func GammaSkew(k, λ float64) float64 {
	return 2 / math.Sqrt(k)
}

// GammaReparamAlphaBeta returns the parameters k, λ (shape, scale) of the Gamma distribution calculated from desired mean and standard deviation. 
// α = shape, β = rate
// To be used to reparametrize the Gamma distribution. 
func GammaReparamAlphaBeta(α, β float64) (k, λ float64) {
	k = α
	λ = 1 / β
	return
}

// GammaReparamModeStd returns the parameters k, λ (shape, scale) of the Gamma distribution calculated from mode and standard deviation. 
// It is more intuitive to start with the mode and standard deviation, instead of the mean and standard deviation as used in the Kruschke (2011) book. 
// The reason is that the gamma distribution is typically very skewed, and therefore the location of the mean is not very intuitive. 
// This function computes the shape and rate parameters of the gamma distribution from a desired mode and standard deviation.
// After http://doingbayesiandataanalysis.blogspot.com/2012/01/parameterizing-gamma-distribution-by.html
func GammaReparamModeStd(mode, sd float64) (k, λ float64) {
	β := (mode + math.Sqrt(mode*mode+4*sd*sd)) / (2 * sd * sd)
	k = 1 + mode*β
	λ = 1 / β
	return
}

// GammaReparamMeanStd returns the parameters k, λ (shape, scale) of the Gamma distribution calculated from mean and standard deviation. 
func GammaReparamMeanStd(mean, sd float64) (k, λ float64) {
	/*
		mean =k*λ
		sd*sd = k*λ*λ
		k=sd*sd /(λ*λ)
		k=mean/λ
		sd*sd /(λ*λ)=mean/λ
		sd*sd *λ/(λ*λ)=mean
		λ/(λ*λ)=mean/(sd*sd )
		λ=(sd*sd )/mean
	*/
	λ = (sd * sd) / mean
	k = mean / λ
	return
}

/************** some non-working code

// GammaCDF returns the CDF of the Gamma distribution. // TO BE REIMPLEMENTED
// Analytic solution, did not pass some tests!
func GammaCDF(k float64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || λ < 0 {
			panic(fmt.Sprintf("k < 0 || λ < 0"))
		}
		if x < 0 {
			return 0
		}
		return Iγ(k, x/λ) / Γ(k)
	}
}

// GammaCDFint returns the CDF of the Gamma distribution, for integer k only. 
// Cumulative distribution function, for integer k only
func GammaCDFint(k int64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || λ < 0 {
			panic(fmt.Sprintf("k < 0 || λ < 0"))
		}
		if x < 0 {
			return 0
		}
		return Iγint(k, x/λ) / Γ(float64(k))
	}
}

// Cumulative distribution function, using gamma incomplete integral  DOES NOT WORK !!!
func GammaCDF(k float64, λ float64) func(x float64) float64 {
	return func(x float64) float64 {
		if k < 0 || λ < 0 {
			panic(fmt.Sprintf("k < 0 || λ < 0"))
		}
		if x < 0 {
			return 0
		}
		return IGam(λ, k*x)
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

**************/
