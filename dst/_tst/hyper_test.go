package dst

import (
	"fmt"
	"testing"
)

func TestHypergeomQtl(t *testing.T) {
	var nN, m, nn int64
	nN = 20
	m = 11
	nn = 5
	p := []float64{0.1, 0.3, 0.5, 0.7, 0.9}
	qtl := []float64{2, 2, 3, 3, 4}
	fmt.Println("TestHypergeomQtl")
	for i := 0; i < int(len(p)); i++ {
		q := HypergeometricQtlFor(nN, m, nn, p[i])
		if q != qtl[i] {
			t.Error()
			fmt.Println(p[i], q, qtl[i])
		}
	}
}
