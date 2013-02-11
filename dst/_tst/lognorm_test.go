// test of LogNormal distribution
package dst

import (
	"fmt"
//	"math/rand"
	"testing"
)

// test against known values
func TestLogNormal(t *testing.T) {
	fmt.Println("test of LogNormal distribution: PDF")
	fn := LogNormalPDF(4.2, 1.33)
	x := fn(3.35)
	y := 0.007141318
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of LogNormal distribution: CDF")
	fn = LogNormalCDF(4.2, 1.33)
	x = fn(3.35)
	y = 0.01225937
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of LogNormal distribution: Qtl")
	qtl:= LogNormalQtl(4.2, 1.33)
	x = qtl(0.01225937)
	y = 3.350001
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}
