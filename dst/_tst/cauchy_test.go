// test of Cauchy distribution against R:actuar:dinvgamma()
package dst

import (
	"fmt"
	"testing"
)

// test against known values
func TestCauchy(t *testing.T) {
	fmt.Println("test of Cauchy distribution: PDF")
	fn := CauchyPDF(2.2, 1.33)
	x := fn(3.35)
	y := 0.1369451
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Cauchy distribution: CDF")
	fn = CauchyCDF(2.2, 1.33)
	x = fn(3.35)
	y = 0.7269373
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Cauchy distribution: Qtl")
	qtl := CauchyQtl(2.2, 1.33)
	x = qtl(0.7269373)
	y = 3.35
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Cauchy distribution: LnPDF")
	fn = CauchyLnPDF(2.2, 1.33)
	x = fn(3.35)
	y = -1.988175
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Cauchy distribution: LnCDF")
	fn = CauchyLnCDF(2.2, 1.33)
	x = fn(3.35)
	y = -0.318915
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}
