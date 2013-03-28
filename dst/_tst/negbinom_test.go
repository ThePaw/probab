package dst

import (
	"fmt"
	"testing"
)

// test against known values
func TestNegativeBinomial_PMF_CDF(t *testing.T) {
	var (
		ρ    float64
		i, n int64
	)

	// edit the following values:  >>>
	ρ = 0.5
	n = 20

	k := []int64{10, 11, 12, 16, 25, 40}
	pmf := []float64{0.0186544004827737808228, 0.025437818840146064758, 0.0328571826685220003128, 0.05907974191359244287, 0.04004139896255765052, 0.00121194851197753156874}
	cdf := []float64{0.0493685733526945114136, 0.074806392192840576172, 0.1076635748613625764847, 0.30885965851484797895, 0.81435098276449480181, 0.9968911986703366647292}

	// <<<

	fmt.Println("test of NegativeBinomial PMF")
	for i = 0; i < int64(len(k)); i++ {
		prob := NegBinomialPMFAt(ρ, n, k[i])
		if !check(prob, pmf[i]) {
			t.Error()
			fmt.Println(k[i], prob, pmf[i])

		}
	}

	fmt.Println("test of NegativeBinomial CDF")
	for i = 0; i < int64(len(k)); i++ {
		prob := NegBinomialCDFAt(ρ, n, k[i])
		if !check(prob, cdf[i]) {
			t.Error()
			fmt.Println(k[i], prob, cdf[i])
		}
	}
}

func TestNegativeBinomial_Qtl(t *testing.T) {
	var (
		ρ    float64
		i, n int64
	)
	ρ = 0.5
	n = 20
	p := []float64{0.1, 0.3, 0.5, 0.7, 0.9}
	qtl := []int64{12, 16, 19, 23, 28}
	fmt.Println("test of NegativeBinomial Qtl")
	for i = 0; i < int64(len(p)); i++ {
		q := NegBinomialQtlFor(ρ, n, p[i])
		if q != qtl[i] {
			t.Error()
			fmt.Println(p[i], q, qtl[i])
		}
	}
}
