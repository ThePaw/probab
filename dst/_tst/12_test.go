// test of Pareto-II distribution
package dst

import (
	"fmt"
	"math"
	"testing"
)

// test against known values: PDF
func TestParetoII_PDF(t *testing.T) {
	θ := 4.0
	α := 3.0
	v := 2.0
	fmt.Println("test of Pareto-II distribution: PDF")
	fn := ParetoII_PDF(θ, α)
	x := fn(v)
	y := 0.1481481
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	θ = 1.0567
	α = 2.1865
	v = 8.7896
	fn = ParetoII_PDF(θ, α)
	x = fn(v)
	y = 0.001686768
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

// test against known values: CDF
func TestParetoII_CDF(t *testing.T) {
	fmt.Println("test of Pareto-II distribution: CDF")
	θ := 1.0567
	α := 2.1865
	v := 8.7896
	fn := ParetoII_CDF(θ, α)
	x := fn(v)
	y := 0.9924041
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

// test against known values: Qtl
func TestParetoII_Qtl(t *testing.T) {
	fmt.Println("test of Pareto-II distribution: Qtl")
	θ := 1.0567
	α := 2.1865
	v := 0.6543
	fn := ParetoII_Qtl(θ, α)
	x := fn(v)
	y := 0.6609253
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestParetoIIMean(t *testing.T) {
	fmt.Println("test of Pareto-II distribution: Mean")
	θ := 1.0567
	α := 2.1865
	x := ParetoII_Mean(θ, α)
	y := 0.8906026
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestParetoIIVar(t *testing.T) {
	fmt.Println("test of Pareto-II distribution: Variance")
	θ := 1.0567
	α := 2.1865
	x := ParetoII_Var(θ, α)
	y := 10.09222
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestParetoIISkew(t *testing.T) {
	fmt.Println("test of Pareto-II distribution: Skewness")
	θ := 7.0
	α := 8.0
	x := ParetoII_Skew(θ, α)
	y := 9.8
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestNextParetoII(t *testing.T) {
	fmt.Println("test of Pareto-II distribution: Random number")
	const nIter = 1e7
	const eps = 1e-3
	θ := 1.0567
	α := 2.1865
	x := 0.0
	for i := 0; i < nIter; i++ {
		x += NextParetoII(θ, α)
	}
	x /= nIter // mean
	y := 0.8906026
	if math.Abs(x/y-1) > eps {
		t.Error()
		fmt.Println(x, y)
	}
}
