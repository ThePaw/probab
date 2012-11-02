// test of Zipf-Mandelbrot distribution
package prob

import (
	"fmt"
//	"math/rand"
	"testing"
)

// test against known values
func TestZipfMandelbrot_PDF(t *testing.T) {
	fmt.Println("test of Zipf-Mandelbrot distribution: PDF")
	fn := ZipfMandelbrotPMF(100, 3, 2)
//	fn := ZipfMandelbrotPMF(N, b, s)

	x := fn(5)
	y := 0.05699202
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

// test against known values
func TestZipfMandelbrot_CDF(t *testing.T) {
	fmt.Println("test of Zipf-Mandelbrot distribution: CDF")
	fn := ZipfMandelbrotCDF(100, 3, 2)
	x := fn(5)
	y := 0.6066174
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

// test against known values
func TestZipfMandelbrotQtl(t *testing.T) {
	fmt.Println("test of Zipf-Mandelbrot distribution: Qtl")
	fn := ZipfMandelbrotQtl(100, 3, 2)
	x := float64(fn(0.6))
	y := float64(5)
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

// test against known values
func TestZipfMandelbrotQtl2(t *testing.T) {
	fmt.Println("test of Zipf-Mandelbrot distribution: Qtl")
	fn := ZipfMandelbrotQtl(100, 3, 2)
	x := float64(fn(0.01))
	y := float64(1)
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

// test against known values
func TestNextZipfMandelbrot(t *testing.T) {
	fmt.Println("test of Zipf-Mandelbrot distribution: Random number")
	const nIter = 1e7
	x := 0.0
	for i :=0; i< nIter; i++{
		x += float64(NextZipfMandelbrot(100, 3, 2))
	}
	x /= nIter
	y := float64(9.341125927812211)
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

func TestZipfMandelbrotMean(t *testing.T) {
	fmt.Println("test of Zipf-Mandelbrot distribution: Mean")
	x := ZipfMandelbrotMean(100, 3, 2)
	y := float64(9.341125927812211)
	if !check(x, y) {
		t.Error()
		fmt.Println(x, y)
	}
}

