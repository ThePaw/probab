// Test Bolstad 2007
package bayes

import (
	"fmt"
	"testing"
)


func TestNormalMuDiffPosterior2(t *testing.T) {
	fmt.Println("Normal posterior test against Michelson data: Bolstad:242 known variances")
	

	// func NormalMuDiff_PDF(nObs1, nObs2 int, ȳ1, ȳ2, σ, μ1Pri, σ1Pri, μ2Pri, σ2Pri float64) func(x float64) float64

	pdf := NormalMuDiff_PDF_NPriKn(20, 23 , 299909, 299756.217391304, 100, 100, 300000, 500, 300000, 500)
	alpha:= []float64{100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}
	known := []float64{0.002975008,0.004951945,0.007404859,0.009947435,0.0120049,0.01301546,0.01267692,0.01109229,0.008719312,0.006157386,0.003906284}

	// check
	n := len(known)
	for i := 0; i < n; i++ {
		x := pdf(alpha[i])
		y := known[i]
		if !check(x, y) {
			t.Error()
			fmt.Println(i, x, y)
		}
	}
}

