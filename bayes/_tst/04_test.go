//
package bayes

import (
	"fmt"
	"testing"
)

// Normal posterior test against Bolstad:normnp
func TestNormalPosterior2(t *testing.T) {
	fmt.Println("Normal posterior test against Bolstad:201")
	

	// func NormMuSingle_PMF_DPri(y, sigma float64, mu, prior []float64) (post []float64)
	known := []float64{0.1238,0.1991,0.2493,0.2431,0.1847}
	mu := []float64{2.0,2.5,3.0,3.5,4.0}
	prior := []float64{0.2,0.2,0.2,0.2,0.2}

	// check
	n := len(mu)
	post := NormMuSingle_PMF_DPri(3.2, 1 , mu, prior)
	for i := 0; i < n; i++ {
		x := post[i]
		y := known[i]
		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

func TestNormalPosterior3(t *testing.T) {
	fmt.Println("Normal posterior test against Bolstad:203")
	// func NormMu_PMF_DPri(nObs int, samp_mean, sigma float64, mu []float64, prior []float64) (post []float64)
	known := []float64{3.855124e-01,4.561038e-01,1.546041e-01,3.753619e-03,2.611027e-05}
	mu := []float64{2.0,2.5,3.0,3.5,4.0}
	prior := []float64{0.1,0.2,0.4,0.2,0.1}

	// check
	n := len(mu)
	post := NormMu_PMF_DPri(5, 2.04, 1, mu, prior)
	for i := 0; i < n; i++ {
		x := post[i]
		y := known[i]
		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

func TestNormalPosterior4(t *testing.T) {
	fmt.Println("Normal posterior test against Bolstad:210, Flat prior")
	// func NormMu_Qtl_FPri(nObs int, ȳ, σ, p float64) float64 {
	alpha:= []float64{0.005, 0.01, 0.025, 0.05, 0.5, 0.95, 0.975, 0.99, 0.995}
	known := [...]float64{30.51272,30.65677,30.86832,31.05026,32,32.94974,33.13168,33.34323,33.48728}

	// check
	for i := 0; i < 9; i++ {
		x := NormMu_Qtl_FPri(12, 32, 2, alpha[i])
		y := known[i]
		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}


