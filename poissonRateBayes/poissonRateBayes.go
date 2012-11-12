// Summary of the posterior distribution of the Poisson parameter. 

package main

import (
	"fmt"
	"code.google.com/p/probab/bayes"
)

// Summary of the posterior distribution of the Poisson parameter. 
func main() {
	var (
		x, n int64
		r, v float64
)

	fmt.Scanf("%d %d %f %f", x, n, r, v, )

	pr := []float64{0.005,0.01,0.025,0.05,0.5,0.95,0.975,0.99,0.995}

	if r<0 || v<0 {
		panic("Shape parameter r and rate parameter v must be greater than or equal to zero")
	}
	fmt.Println("\nProb.\t\tQuantile \n")
	for i:=0; i< 9; i++ {
			qf := bayes.PoissonLambdaQtlGPri(x, n, r, v)
			qtl := qf(pr[i])
			fmt.Println(pr[i],"\t\t", qtl)
	}
	fmt.Println("\n")
}


