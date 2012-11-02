// test of Planck distribution
package prob

import (
	"fmt"
//	"math/rand"
	"testing"
)
 NextPlanck(a, b float64) (x float64)
func TestNextPlanck(t *testing.T) {
	fmt.Println("test NextPlanck")
	const nIter = 1e5
	x := 0.0
	y := PlanckMean(2,3)
	for i :=0; i< nIter; i++{
		x += float64(NextPlanck(2,3))
	}
	x /= nIter
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}


