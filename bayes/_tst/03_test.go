//
package bayes

import (
	"fmt"
	"testing"
)

// Normal posterior test against Bolstad:normnp
func TestNormalPosterior(t *testing.T) {
	fmt.Println("Normal posterior test against Bolstad:normnp")

	// func NormMuPostMean(n int, samp_mean, sigma, pri_mu, pri_sigma float64) float64
	x := NormMuPostMean(20, -0.6267289, 1, 0, 1)
	y := -0.5968847
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMuPostMean: ", x, y)
	}

	x = NormMuPostStd(20, 1, 0, 1)
	y = 0.2182179
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMuPostStd: ", x, y)
	}

	// func NormMu_Qtl_NPri(n int, samp_mean, sigma, pri_mu, pri_sigma, p float64) float64
	x = NormMu_Qtl_NPri(20, -0.6267289, 1, 0, 1, 0.005)
	y = -1.1589767
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMu_Qtl_NPri 0.005: ", x, y)
	}

	x = NormMu_Qtl_NPri(20, -0.6267289, 1, 0, 1, 0.025)
	y = -1.0245839
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMu_Qtl_NPri 0.025: ", x, y)
	}

	x = NormMu_Qtl_NPri(20, -0.6267289, 1, 0, 1, 0.95)
	y = -0.2379482
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMu_Qtl_NPri 0.025: ", x, y)
	}

	// func NormMuPostMean(n int, samp_mean, sigma, pri_mu, pri_sigma float64) float64
	x = NormMuPostMean(20, 0.8091967, 1.33, 0.512, 3.05)
	y = 0.8063977
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMuPostMean: ", x, y)
	}

	x = NormMuPostStd(20, 1.33, 0.512, 3.05)
	y = 0.2959933
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMuPostStd: ", x, y)
	}

	// func NormMu_Qtl_NPri(n int, samp_mean, sigma, pri_mu, pri_sigma, p float64) float64

	x = NormMu_Qtl_NPri(20, 0.8091967, 1.33, 0.512, 3.05, 0.95)
	y = 1.2932633
	if !check(x, y) {
		t.Error()
		fmt.Println("NormMu_Qtl_NPri 0.95: ", x, y)
	}

	alpha := []float64{0.005, 0.01, 0.025, 0.05, 0.5, 0.95, 0.975, 0.99, 0.995}
	known := [...]float64{0.0439696, 0.1178144, 0.2262616, 0.3195321, 0.8063977, 1.2932633, 1.3865338, 1.494981, 1.5688258}

	// check
	for i := 0; i < 9; i++ {
		x = NormMu_Qtl_NPri(20, 0.8091967, 1.33, 0.512, 3.05, alpha[i])
		y = known[i]
		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}

	/*
		// func NormMu_Qtl_FPri(n int, samp_mean, sigma, p float64) float64
		known = [...]float64{-0.2970927,-0.2014596,-0.0610146,0.0597758,0.6902934,1.320811,1.4416014,1.5820465,1.6776795}

		// check
		for i := 0; i < 9; i++ {
			x = NormMu_Qtl_FPri(20, 0.8091967, 1, alpha[i])
			y = known[i]
			if !check(x, y) {
				t.Error()
				fmt.Println(i, x, y)
			}
		}
	*/

	// func NormMu_CrI_NPriKnown(n int, samp_mean, sigma, pri_mu, pri_sigma, alpha float64) (lo, hi float64)
	// check
	x, _ = NormMu_CrI_NPriKnown(20, 0.8091967, 1.33, 0.512, 3.05, 0.05)
	y = 0.2262616
	if !check(x, y) {
		t.Error()
		fmt.Println("lower: ", x, y)
	}
	_, x = NormMu_CrI_NPriKnown(20, 0.8091967, 1.33, 0.512, 3.05, 0.05)
	y = 1.3865338
	if !check(x, y) {
		t.Error()
		fmt.Println("upper: ", x, y)
	}

	// func NormMuSingle_Qtl_NPri(y, sigma, pri_mu, pri_sigma, p float64) float64
	known = [...]float64{-3.2627996, -2.9586492, -2.5119794, -2.1278188, -0.1225275, 1.8827638, 2.2669245, 2.7135942, 3.0177447}

	// check
	for i := 0; i < 9; i++ {
		x = NormMuSingle_Qtl_NPri(-0.2431849, 1.33, 0.512, 3.05, alpha[i])
		y = known[i]
		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}
