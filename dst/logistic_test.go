// test of Logistic distribution
package dst

import (
	"fmt"
	//	"math/rand"
	"testing"
)

// test against known values
func TestLogistic(t *testing.T) {
	fmt.Println("test of Logistic distribution: PDF")
	fn := LogisticPDF(4.2, 1.33)
	x := fn(3.35)
	y := 0.1700109
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Logistic distribution: CDF")
	fn = LogisticCDF(4.2, 1.33)
	x = fn(3.35)
	y = 0.3454505
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Logistic distribution: Qtl")
	qtl := LogisticQtl(4.2, 1.33)
	x = qtl(0.3454505)
	y = 3.35
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}
