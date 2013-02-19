package dst

import (
	"math"
)

const π = float64(math.Pi)
const Ln2 = math.Ln2
const M_LN_SQRT_2PI = 0.918938533204672741780329736406 // log(sqrt(2*pi))
const min64 = math.SmallestNonzeroFloat64              //   DBL_MIN
const eps64 = 1.1102230246251565e-16                   // DBL_EPSILON   
const maxExp = 1024.0                                  // DBL_MAX_EXP

var NaN = math.NaN()

var fZero float64 = float64(0.0)
var fOne float64 = float64(1.0)
var iZero int64 = int64(0)
var iOne int64 = int64(1)

var negInf float64 = math.Inf(-1)
var posInf float64 = math.Inf(+1)

var abs func(float64) float64 = math.Abs
var floor func(float64) float64 = math.Floor
var log func(float64) float64 = math.Log
var log1p func(float64) float64 = math.Log1p
var exp func(float64) float64 = math.Exp
var sqrt func(float64) float64 = math.Sqrt
var pow func(float64, float64) float64 = math.Pow
var atan func(float64) float64 = math.Atan
var tan func(float64) float64 = math.Tan
var trunc func(float64) float64 = math.Trunc

var isNaN func(float64) bool = math.IsNaN
var isInf func(float64, int) bool = math.IsInf

func imin(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

func imax(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func min(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func maxFloat64(x []float64) float64 {
	first := x[0]
	if len(x) > 1 {
		rest := maxFloat64(x[1:len(x)])
		if rest > first {
			first = rest
		}
	}
	return first
}

func maxInt64(x []int64) int64 {
	first := x[0]
	if len(x) > 1 {
		rest := maxInt64(x[1:len(x)])
		if rest > first {
			first = rest
		}
	}
	return first
}

func copyInt64(x []int64, n int64) []int64 {
	newx := make([]int64, n)
	for i := 0; i < len(x) && i < int(n); i++ {
		newx[i] = x[i]
	}
	return newx
}

func copyFloat64(x []float64, n int64) []float64 {
	newx := make([]float64, n)
	for i := 0; i < len(x) && i < int(n); i++ {
		newx[i] = x[i]
	}
	return newx
}

func expm1(x float64) float64 {
	var y float64
	a := math.Abs(x)
	if a < math.SmallestNonzeroFloat64 {
		y = x
	} else if a > 0.697 {
		y = exp(x) - 1 /* negligible cancellation */
	} else {
		if a > 1e-8 {
			y = exp(x) - 1
		} else { /* Taylor expansion, more accurate in this range */
			y = (x/2 + 1) * x
		}
		/* Newton step for solving   log(1 + y) = x   for y : */
		/* WARNING: does not work for y ~ -1: bug in 1.5.0 */
		y -= (1 + y) * (math.Log1p(y) - x)
	} //else
	return y
}

func RejectionSample(targetDensity func(float64) float64, sourceDensity func(float64) float64, source func() float64, K float64) float64 {
	x := source()
	for ; UniformNext(0, 1) >= targetDensity(x)/(K*sourceDensity(x)); x = source() {

	}
	return x
}

func ShuffleInt64(x []int64) {
	n := int64(len(x))
	for i := iZero; i < n; i++ {
		j := i + RangeNext(n-i)
		t := x[i]
		x[i] = x[j]
		x[j] = t
	}
}

func ShuffleFloat64(x []float64) {
	n := int64(len(x))
	for i := iZero; i < n; i++ {
		j := i + RangeNext(n-i)
		t := x[i]
		x[i] = x[j]
		x[j] = t
	}
}

func Shuffle(x []interface{}) {
	n := int64(len(x))
	for i := iZero; i < n; i++ {
		j := i + RangeNext(n-i)
		t := x[i]
		x[i] = x[j]
		x[j] = t
	}
}

func iround(x float64) float64 {
	return math.Floor(x + 0.5) // ?correct -- round to integer
}
