// test of Student's t distribution
package prob

import (
	"fmt"
	"math/rand"
	"testing"
)

// test against known values
func TestStudentsT_PDF(t *testing.T) {
	fmt.Println("test of Student's t distribution: PDF")
	fn := StudentsT_PDF(3)
	x := fn(3)
	y := 0.02297204
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_CDF(t *testing.T) {
	fmt.Println("test of Student's t distribution: CDF")
	fn := StudentsT_CDF(3)
	x := fn(3)
	y := 0.9711656
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_CDF2(t *testing.T) {
	fmt.Println("test of Student's t distribution: CDF #2")
	fn := StudentsT_CDF(99)
	x := fn(3)
	y := 0.9982922
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_CDF3(t *testing.T) {
	fmt.Println("test of Student's t distribution: CDF #3")
	fn := StudentsT_CDF(1)
	x := fn(2e100)
	y := 1.0
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_Qtl(t *testing.T) {
	fmt.Println("test of Student's t distribution: Quantile")
	fn := StudentsT_Qtl(1e25)
	x := fn(0.9982922)
	y := 2.927627
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_Qtl2(t *testing.T) {
	fmt.Println("test of Student's t distribution: Quantile #2")
	fn := StudentsT_Qtl(99)
	x := fn(0.9982922)
	y := 2.999991
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_Qtl3(t *testing.T) {
	fmt.Println("test of Student's t distribution: Quantile #3")
	fn := StudentsT_Qtl(3)
	x := fn(0.5)
	y := 0.0
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestStudentsT_Qtl4(t *testing.T) {
	fmt.Println("test of Student's t distribution: Quantile against CDF, 100000 iterations")

	for i := 0; i < 100000; i++ {
		//func (r *Rand) Intn(n int) int {
		df := float64(rand.Intn(10000) + 1)
		p := rand.Float64()
		fn := StudentsT_Qtl(df)
		x := fn(p)
		cdf := StudentsT_CDF(df)
		y := cdf(x)
		if !check(y, p) {
			t.Error()
			fmt.Println(y, p)
		}
	}
}
