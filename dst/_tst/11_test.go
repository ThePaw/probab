// test of Zeta distribution
package dst

import (
	"fmt"
	//	"math/rand"
	"testing"
)

// test of Zeta distribution

func TestNextZeta(t *testing.T) {
	fmt.Println("test NextZeta")
	const nIter = 1e7
	x := 0.0
	y := ZetaMean(2.152)
	for i := 0; i < nIter; i++ {
		x += float64(NextZeta(2.152))
	}
	x /= nIter
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}
