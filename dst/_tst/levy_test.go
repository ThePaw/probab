// test of Levy distribution against R:rmutil:dlevy()
package dst

import (
	"fmt"
	//	"math/rand"
	"testing"
)

// test against known values
func TestLevy(t *testing.T) {
	fmt.Println("test of Levy distribution: PDF")
	fn := LevyPDF(2.2, 1.33)
	x := fn(3.35)
	y := 0.2092443
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Levy distribution: CDF")
	fn = LevyCDF(2.2, 1.33)
	x = fn(3.35)
	y = 0.2821881
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Levy distribution: Qtl")
	qtl := LevyQtl(2.2, 1.33)
	x = qtl(0.2821881)
	y = 3.35
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of Levy distribution: LnPDF")
	fn = LevyLnPDF(2.2, 1.33)
	x = fn(3.35)
	y = -1.564253
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}
