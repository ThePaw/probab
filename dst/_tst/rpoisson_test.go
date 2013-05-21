package dst

import (
	"fmt"
	"testing"
)

func TestPoissonNext2(t *testing.T) {
	var λ float64
	λ = 1000
	n := 0.0
	mean := 0.0
	m2 := 0.0
	//	for j := 0; j < 100000000; j++ {
	for j := 0; j < 1000; j++ {
		n++
		x := float64(PoissonNext(λ))
		delta := x - mean
		mean += delta / n
		if n > 1 {
			m2 += delta * (x - mean)
		}
	}
	v := m2 / n
	fmt.Println("mean, var: ", mean, v)
}
