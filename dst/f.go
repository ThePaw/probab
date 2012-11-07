// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// F-distribution, alias Fisher-Snedecor distribution

import (
	"fmt"
	"math"
	. "code.google.com/p/go-fn/fn"
)

// F_PDF returns the PDF of the F distribution. 
func F_PDF(d1, d2 int64) func(x float64) float64 {
	df1 := float64(d1)
	df2 := float64(d2)
	normalization := 1 / B(df1/2, df2/2)
	return func(x float64) float64 {
		return normalization * sqrt(pow(df1*x, df1)*pow(df2, df2)/pow(df1*x+df2, df1+df2)) / x
	}
}

// F_LnPDF returns the natural logarithm of the PDF of the F distribution. 
func F_LnPDF(d1, d2 int64) func(x float64) float64 {
	df1 := float64(d1)
	df2 := float64(d2)
	normalization := -LnB(df1/2, df2/2)
	return func(x float64) float64 {
		return normalization + log(df1*x)*df1/2 + log(df2)*df2/2 - log(df1*x+df2)*(df1+df2)/2 - log(x)
	}
}

// F_PDF_At returns the value of PDF of F distribution at x. 
func F_PDF_At(d1, d2 int64, x float64) float64 {
	pdf :=F_PDF(d1, d2)
	return pdf(x)
}

// F_CDF returns the CDF of the F distribution. 
func F_CDF(d1, d2 int64) func(x float64) float64 {
	return func(x float64) float64 {
		df1 := float64(d1)
		df2 := float64(d2)
		y := df1 * x / (df1*x + df2)
		return BetaIncReg(df1/2.0, df2/2.0, y)
	}
}


// F_CDF_At returns the value of CDF of the F distribution, at x. 
func F_CDF_At(d1, d2 int64, x float64) float64 {
	cdf:=F_CDF(d1, d2)
	return cdf(x)
}

// F_Qtl returns the inverse of the CDF (quantile) of the F distribution. 
func F_Qtl(d1, d2 int64) func(p float64) float64 {
	df1 := float64(d1)
	df2 := float64(d2)
	return func(p float64) float64 {
	if p < 0.0 {
		panic(fmt.Sprintf("p < 0"))
	}
	if p > 1.0 {
		panic(fmt.Sprintf("p > 1.0"))
	}
	if df1 < 1.0 {
		panic(fmt.Sprintf("df1 < 1"))
	}
	if df2 < 1.0 {
		panic(fmt.Sprintf("df2 < 1"))
	}

	return ((1 / Beta_Qtl_For(df2 / 2, df1 / 2, 1 - p) - 1) * df2 / df1);
	}
}

// F_Qtl_For returns the inverse of the CDF (quantile) of the F distribution, for given probability.
func F_Qtl_For(d1, d2 int64, p float64) float64 {
	cdf :=F_Qtl(d1, d2)
	return cdf(p)
}

// NextF returns random number drawn from the F distribution. 
func NextF(d1, d2 int64) float64 {
	df1 := float64(d1)
	df2 := float64(d2)
	return NextChiSquare(d1) * df2 / (NextChiSquare(d2) * df1)
}

// F returns the random number generator with  F distribution. 
func F(d1, d2 int64) func() float64 {
	return func() float64 {
		return NextF(d1, d2)
	}
}

// FMean returns the mean of the F distribution. 
func FMean(d1, d2 int64) float64 {
	if d2<= 2 {
		panic("mean not defined for df2 <=2")
	}
	df2 := float64(d2)
	return df2/(df2-2)
}

// FMode returns the mode of the F distribution. 
func FMode(d1, d2 int64) float64 {
	if d1<= 2 {
		panic("mode not defined for df1 <=2")
	}
	df1 := float64(d1)
	df2 := float64(d2)
	return ((df1-2)/df1)* (df2/(df2+2))
}

// FVar returns the variance of the F distribution. 
func FVar(d1, d2 int64) float64 {
	if d2<= 4 {
		panic("variance not defined for d2 <= 4")
	}
	df1 := float64(d1)
	df2 := float64(d2)
	return 2*df2*df2*(df1+df2-2) / (df1*(df2-2)*(df2-2)*(df2-4))
}

// FStd returns the standard deviation of the F distribution. 
func FStd(d1, d2 int64) float64 {
	if d2<= 4 {
		panic("standard deviation not defined for d2 <= 4")
	}
	df1 := float64(d1)
	df2 := float64(d2)
	v := 2*df2*df2*(df1+df2-2) / (df1*(df2-2)*(df2-2)*(df2-4))
	return math.Sqrt(v)
}

// FSkew returns the skewness of the F distribution. 
func FSkew(d1, d2 int64) float64 {
	if d2<= 6 {
		panic("skewness not defined for d2 <= 6")
	}
	df1 := float64(d1)
	df2 := float64(d2)
	return (2*df1+df2-2) *math.Sqrt(8*(df2-4)) / (df2-6) *math.Sqrt(df1*(df1+df2-2))
}

// FExKurt returns the excess kurtosis of the F distribution. 
func FExKurt(d1, d2 int64) float64 {
	if d2 <= 8 {
		panic("skewness not defined for d2 <= 8 ")
	}
	df1 := float64(d1)
	df2 := float64(d2)
	return 12 *(df1*(5*df2-22)*(df1+df2-2)+(df2-4)*(df2-2)*(df2-2)) / (df1*(df2-6)*(df2-8)*(df1+df2-2))
}

