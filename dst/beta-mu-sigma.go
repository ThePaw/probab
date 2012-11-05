// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Beta distribution reparametrized using mean and standard deviation. 

// Betaμσ_PDF returns the PDF of the Beta distribution reparametrized using mean and standard deviation. 
func Betaμσ_PDF(μ, σ float64) func(x float64) float64 {
	α := μ*(μ*(1-μ)/(σ*σ)-1)
	β := (1-μ)*(μ*(1-μ)/(σ*σ)-1)
	if σ*σ >= μ*(1-μ) {
		panic("σ too big, α, β out of range")
	}
	return Beta_PDF(α, β)
}

// Betaμσ_LnPDF returns the natural logarithm of the PDF of the Beta distribution reparametrized using mean and standard deviation. 
func Betaμσ_LnPDF(μ, σ float64) func(x float64) float64 {
	α := μ*(μ*(1-μ)/(σ*σ)-1)
	β := (1-μ)*(μ*(1-μ)/(σ*σ)-1)
	if σ*σ >= μ*(1-μ) {
		panic("σ too big, α, β out of range")
	}
	return Beta_LnPDF(α, β)
}

// NextBetaμσ returns random number drawn from the  Beta distribution reparametrized using mean and standard deviation. 
func NextBetaμσ(μ, σ float64) float64 {
	α := μ*(μ*(1-μ)/(σ*σ)-1)
	β := (1-μ)*(μ*(1-μ)/(σ*σ)-1)
	if σ*σ >= μ*(1-μ) {
		panic("σ too big, α, β out of range")
	}
	return NextBeta(α, β)
}

// Betaμσ returns the random number generator with  Beta distribution reparametrized using mean and standard deviation. 
func Betaμσ(μ, σ float64) func() float64 {
	α := μ*(μ*(1-μ)/(σ*σ)-1)
	β := (1-μ)*(μ*(1-μ)/(σ*σ)-1)
	if σ*σ >= μ*(1-μ) {
		panic("σ too big, α, β out of range")
	}
	return func() float64 { return NextBeta(α, β) }
}

// Betaμσ_PDF_At returns the value of PDF of Beta distribution(μ, σ) at x. 
func Betaμσ_PDF_At(μ, σ, x float64) float64 {
	pdf := Betaμσ_PDF(μ, σ)
	return pdf(x)
}

// Betaμσ_CDF returns the CDF of the Beta distribution reparametrized using mean and standard deviation. 
func Betaμσ_CDF(μ, σ float64) func(x float64) float64 {
	α := μ*(μ*(1-μ)/(σ*σ)-1)
	β := (1-μ)*(μ*(1-μ)/(σ*σ)-1)
	if σ*σ >= μ*(1-μ) {
		panic("σ too big, α, β out of range")
	}
	return Beta_CDF(α , β)
}

// Betaμσ_CDF_At returns the value of CDF of the Beta distribution reparametrized using mean and standard deviation, at x. 
func Betaμσ_CDF_At(μ, σ, x float64) float64 {
	cdf := Beta_CDF(μ, σ)
	return cdf(x)
}

// Betaμσ_Qtl returns the inverse of the CDF (quantile) of the Beta distribution reparametrized using mean and standard deviation. 
func Betaμσ_Qtl(μ, σ float64) func(p float64) float64 {
	// p: probability for which the quantile is evaluated
	α := μ*(μ*(1-μ)/(σ*σ)-1)
	β := (1-μ)*(μ*(1-μ)/(σ*σ)-1)
	if σ*σ >= μ*(1-μ) {
		panic("σ too big, α, β out of range")
	}
	return Beta_Qtl(α, β)
}

// Betaμσ_Qtl_For returns the inverse of the CDF (quantile) of the Beta distribution reparametrized using mean and standard deviation, for a given probability.
func Betaμσ_Qtl_For(μ, σ, p float64) float64 {
	cdf := Betaμσ_Qtl(μ, σ)
	return cdf(p)
}

