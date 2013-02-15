// test of InvGamma distribution against R:actuar:dinvgamma()
package dst

import (
	"fmt"
	//	"math/rand"
	"testing"
)

// test against known values
func TestInvGamma(t *testing.T) {
	fmt.Println("test of InvGamma distribution: PDF")
	fn := InvGammaPDF(2.2, 1.33)
	x := fn(3.35)
	y := 0.0238674
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of InvGamma distribution: CDF")
	fn = InvGammaCDF(2.2, 1.33)
	x = fn(3.35)
	y = 0.9586864
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of InvGamma distribution: Qtl")
	qtl := InvGammaQtl(2.2, 1.33)
	x = qtl(0.9586864)
	y = 3.35

	fmt.Println("GammaQtlFor", GammaQtlFor(2.2, 1.0, 1-0.9586864))
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}

	fmt.Println("test of InvGamma distribution: LnPDF")
	fn = InvGammaLnPDF(2.2, 1.33)
	x = fn(3.35)
	y = -3.735242
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}
