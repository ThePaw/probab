// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Beta distribution reparametrized using mean (μ) and sample size (ν). 
// Kruschke, J. K. (2011). Doing Bayesian data analysis: A tutorial with R and BUGS. p. 83: Academic Press / Elsevier. ISBN 978-0123814852.

// Betaμν_PDF returns the PDF of the Beta distribution reparametrized using mean and sample size. 
func Betaμν_PDF(μ, ν float64) func(x float64) float64 {
	if ν <= 0 {
		panic("ν must be greater than zero")
	}
	α := μ*ν
	β := (1-μ)*ν
	return Beta_PDF(α, β)
}

// Betaμν_LnPDF returns the natural logarithm of the PDF of the Beta distribution reparametrized using mean and sample size. 
func Betaμν_LnPDF(μ, ν float64) func(x float64) float64 {
	α := μ*ν
	β := (1-μ)*ν
	if ν <= 0 {
		panic("ν must be greater than zero")
	}
	return Beta_LnPDF(α, β)
}

// NextBetaμν returns random number drawn from the  Beta distribution reparametrized using mean and sample size. 
func NextBetaμν(μ, ν float64) float64 {
	α := μ*ν
	β := (1-μ)*ν
	if ν <= 0 {
		panic("ν must be greater than zero")
	}
	return NextBeta(α, β)
}

// Betaμν returns the random number generator with  Beta distribution reparametrized using mean and sample size. 
func Betaμν(μ, ν float64) func() float64 {
	α := μ*ν
	β := (1-μ)*ν
	if ν <= 0 {
		panic("ν must be greater than zero")
	}
	return func() float64 { return NextBeta(α, β) }
}

// Betaμν_PDF_At returns the value of PDF of Beta distribution at x. 
func Betaμν_PDF_At(μ, ν, x float64) float64 {
	pdf := Betaμν_PDF(μ, ν)
	return pdf(x)
}

// Betaμν_CDF returns the CDF of the Beta distribution reparametrized using mean and sample size. 
func Betaμν_CDF(μ, ν float64) func(x float64) float64 {
	α := μ*ν
	β := (1-μ)*ν
	if ν <= 0 {
		panic("ν must be greater than zero")
	}
	return Beta_CDF(α , β)
}

// Betaμν_CDF_At returns the value of CDF of the Beta distribution reparametrized using mean and sample size, at x. 
func Betaμν_CDF_At(μ, ν, x float64) float64 {
	cdf := Beta_CDF(μ, ν)
	return cdf(x)
}

// Betaμν_Qtl returns the inverse of the CDF (quantile) of the Beta distribution reparametrized using mean and sample size. 
func Betaμν_Qtl(μ, ν float64) func(p float64) float64 {
	// p: probability for which the quantile is evaluated
	α := μ*ν
	β := (1-μ)*ν
	if ν <= 0 {
		panic("ν must be greater than zero")
	}
	return Beta_Qtl(α, β)
}

// Betaμν_Qtl_For returns the inverse of the CDF (quantile) of the Beta distribution reparametrized using mean and sample size, for a given probability.
func Betaμν_Qtl_For(μ, ν, p float64) float64 {
	cdf := Betaμν_Qtl(μ, ν)
	return cdf(p)
}

