// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Beta distribution. 
// Parameters:
// α > 0: shape
// β > 0: shape. 
// Support: x ∈ [0; 1].


import (
	. "code.google.com/p/go-fn/fn"
	"fmt"
	"math"
)

func bisect(x, p, a, b, xtol, ptol float64) float64 {
	var x0, x1, px float64
	cdf := Beta_PDF(a, b)
	for math.Abs(x1-x0) > xtol {
		px = cdf(x)
		switch {
		case math.Abs(px-p) < ptol:
			return x
		case px < p:
			x0 = x
		case px > p:
			x1 = x
		}
		x = 0.5 * (x0 + x1)
	}
	return x
}

func betaContinuedFraction(α, β, x float64) float64 {
	var aa, del, res, qab, qap, qam, c, d, m2, m, acc float64
	var i int64
	const eps = 2.2204460492503131e-16
	const maxIter = 1000000000

	acc = 1e-16
	qab = α + β
	qap = α + 1.0
	qam = α - 1.0
	c = 1.0
	d = 1.0 - qab*x/qap

	if math.Abs(d) < eps {
		d = eps
	}
	d = 1.0 / d
	res = d

	for i = 1; i <= maxIter; i++ {
		m = (float64)(i)
		m2 = 2 * m
		aa = m * (β - m) * x / ((qam + m2) * (α + m2))
		d = 1.0 + aa*d
		if math.Abs(d) < eps {
			d = eps
		}
		c = 1.0 + aa/c
		if math.Abs(c) < eps {
			c = eps
		}
		d = 1.0 / d
		res *= d * c
		aa = -(α + m) * (qab + m) * x / ((α + m2) * (qap + m2))
		d = 1.0 + aa*d
		if math.Abs(d) < eps {
			d = eps
		}
		c = 1.0 + aa/c
		if math.Abs(c) < eps {
			c = eps
		}
		d = 1.0 / d
		del = d * c
		res *= del
		if math.Abs(del-1.0) < acc {
			return res
		}
	}

	panic(fmt.Sprintf("betaContinuedFraction(): α or β too big, or maxIter too small"))
	return -1.00
}

// Beta_PDF returns the PDF of the Beta distribution. 
func Beta_PDF(α, β float64) func(x float64) float64 {
	if α == 1 && β == 1{	// uniform case
		return Uniform_PDF(0, 1)
	}
	dα := []float64{α, β}
	dirPDF := Dirichlet_PDF(dα)
	return func(x float64) float64 {
		if 0 > x || x > 1 {
			return 0
		}
		dx := []float64{x, 1 - x}
		return dirPDF(dx)
	}
}

// Beta_LnPDF returns the natural logarithm of the PDF of the Beta distribution. 
func Beta_LnPDF(α, β float64) func(x float64) float64 {
	dα := []float64{α, β}
	dirLnPDF := Dirichlet_LnPDF(dα)
	return func(x float64) float64 {
		if 0 > x || x > 1 {
			return negInf
		}
		dx := []float64{x, 1 - x}
		return dirLnPDF(dx)
	}
}

// Beta_PDF_At returns the value of PDF of Beta distribution at x. 
func Beta_PDF_At(α, β, x float64) float64 {
	pdf := Beta_PDF(α, β)
	return pdf(x)
}

// Beta_CDF returns the CDF of the Beta distribution. 
func Beta_CDF(α, β float64) func(x float64) float64 {
	if α == 1 && β == 1{	// uniform case
		return Uniform_CDF(0, 1)
	}
	return func(x float64) float64 {
		var y, res float64
		y = math.Exp(LnΓ(α+β) - LnΓ(α) - LnΓ(β) + α*math.Log(x) + β*math.Log(1.0-x))
		switch {
		case x == 0:
			res = 0.0
		case x == 1.0:
			res = 1.0
		case x < (α+1.0)/(α+β+2.0):
			res = y * betaContinuedFraction(α, β, x) / α
		default:
			res = 1.0 - y*betaContinuedFraction(β, α, 1.0-x)/β

		}
		return res
	}
}

// Beta_CDF_At returns the value of CDF of the Beta distribution, at x. 
func Beta_CDF_At(α, β, x float64) float64 {
	cdf := Beta_CDF(α, β)
	return cdf(x)
}

// Beta_Qtl returns the inverse of the CDF (quantile) of the Beta distribution. 
func Beta_Qtl(α, β float64) func(p float64) float64 {
	// p: probability for which the quantile is evaluated
	return func(p float64) float64 {
		var x float64 = 0
		var a float64 = 0
		var b float64 = 1
		var precision float64 = 1e-9
		if p < 0.0 {
			panic(fmt.Sprintf("p < 0"))
		}
		if p > 1.0 {
			panic(fmt.Sprintf("p > 1.0"))
		}
		if α < 0.0 {
			panic(fmt.Sprintf("α < 0.0"))
		}
		if β < 0.0 {
			panic(fmt.Sprintf("β < 0.0"))
		}

		for (b - a) > precision {
			x = (a + b) / 2
			if BetaIncReg(α, β, x) > p {
				b = x
			} else {
				a = x
			}
		}

		return x
	}
}

// Beta_Qtl_For returns the inverse of the CDF (quantile) of the Beta distribution, for given probability.
func Beta_Qtl_For(α, β, p float64) float64 {
	cdf := Beta_Qtl(α, β)
	return cdf(p)
}

// NextBeta returns random number drawn from the Beta distribution. 
func NextBeta(α, β float64) float64 {
	if α == 1 && β == 1{	// uniform case
		return NextUniform(0, 1)
	}
	dα := []float64{α, β}
	return NextDirichlet(dα)[0]
}

// Beta returns the random number generator with  Beta distribution. 
func Beta(α, β float64) func() float64 {
	if α == 1 && β == 1{	// uniform case
		return Uniform(0, 1)
	}
	return func() float64 { return NextBeta(α, β) }
}

// BetaMean returns the mean of the Beta distribution. 
func BetaMean(α, β float64) (μ float64) {
	if α == β  {		// symmetric case
		μ = 0.5
	} else {
		μ = α / (α + β)
	}
	return 
}

// BetaMedian returns the median of the Beta distribution. 
func BetaMedian(α, β float64) (med float64) {
	//The median of the beta distribution is the unique real number 
	// for which the regularized incomplete beta function  = 0.5 . 
	// There is no general closed-form expression for the median of the beta distribution 
	// for arbitrary values of α and β. 
	switch {
	case α == β: 	// symmetric case
		med = 0.5
	case α == 1 && β >0: 
		med = 1.0 - math.Pow(0.5, 1/β)
	case β == 1 && α >0: 
		med = math.Pow(0.5, 1/α)
	case α == 3 && β == 2: 
		med = 0.6142724318676105
	case α == 2 && β == 3: 
		med = 0.38572756813238945
	case α <= 1 || β <= 1: 
		med = (α - 1/3) / (α + β - 2/3)		// approximation
	default:
		panic("no closed form for median, sorry")
	}
	return
}

// BetaMedianApprox returns the approximate median of the Beta distribution. 
func BetaMedianApprox(α, β float64) float64 {
	if α <= 1 || β <= 1 {
		panic("α<=1 || β<=1")
	}
	return (α - 1/3) / (α + β - 2/3)
}

// BetaMode returns the mode of the Beta distribution. 
func BetaMode(α, β float64) float64 {
	if α <= 1 || β <= 1 {
		panic("α<=1 || β<=1")
	}
	return (α - 1) / (α + β - 2) // if α < 1 and β < 1, this is the anti-mode
}

// BetaVar returns the variance of the Beta distribution. 
func BetaVar(α, β float64) float64 {
	return (α * β) / ((α + β) * (α + β) * (α + β + 1))
}

// BetaStd returns the standard deviation of the Beta distribution. 
func BetaStd(α, β float64) float64 {
	v := (α * β) / ((α + β) * (α + β) * (α + β + 1))
	return math.Sqrt(v)
}

// BetaSkew returns the skewness of the Beta distribution. 
func BetaSkew(α, β float64) (s float64) {

	if α == β {	// symmetric case
		s = 0.0
	} else {
		num := 2 * (β - α) * math.Sqrt(α + β + 1)
		den := (α + β + 2) * math.Sqrt(α * β)
		s = num / den
	}
	return
}

// BetaExKurt returns the excess kurtosis of the Beta distribution. 
func BetaExKurt(α, β float64) float64 {
	num := 6*((α - β)*(α - β)*(α + β + 1) - α * β * (α + β + 2))
	den := α * β * (α + β + 2) * (α + β + 2) 
	return num / den
}

