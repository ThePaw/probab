// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

import (
	"code.google.com/p/probab/dst"
	"math/rand"
)

/*

gibbs=function(logpost,start,m,scale,...)
{ 
p=length(start)
vth=array(0,dim=c(m,p))
f0=logpost(start,...)
arate=array(0,dim=c(1,p))

th0=start
for (i in 1:m)
{
  for (j in 1:p)
  {
  th1=th0
  th1[j]=th0[j]+rnorm(1)*scale[j]
  f1=logpost(th1,...)
  u=runif(1)<exp(f1-f0)
  th0[j]=th1[j]*(u==1)+th0[j]*(u==0)
  f0=f1*(u==1)+f0*(u==0)
  vth[i,j]=th0[j]; 
  arate[j]=arate[j]+u
  }
}
arate=arate/m
stuff=list(par=vth,accept=arate)
return(stuff)
}


*/

// Metropolis within Gibbs sampling algorithm of a posterior distribution.
func Gibbs(logpost func([]float64) float64, start []float64, m int, scale []float64) (vth [][]float64, arate []float64) {
	// Arguments:
	// logpost - function defining the log posterior density
	// start - array with a single row that gives the starting value of the parameter vector
	// m - the number of iterations of the chain
	// scale - vector of scale parameters for the random walk Metropolis steps
	// Returns:
	// par - a matrix of simulated values where each row corresponds to a value of the vector  parameter
	// accept - vector of acceptance rates of the Metropolis steps of the algorithm

	p := len(start)
	//vth=array(0,dim=c(m,p))
	// make vth matrix
	nCol := p
	s := make([]float64, m*nCol)
	vth = make([][]float64, m)
	for i, p := 0, 0; i < m; i++ {
		vth[i] = s[p : p+nCol]
		p += nCol
	}

	f0 := logpost(start)
	//arate=array(0,dim=c(1,p))
	// make arate vector
	arate = make([]float64, p)

	th0 := make([]float64, p)
	for i, val := range start {
		th0[i] = val
	}

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			th1 := make([]float64, p)
			for k, val := range th0 {
				th1[k] = val
			}

			th1[j] = th0[j] + dst.NormalNext(0, 1)*scale[j]
			f1 := logpost(th1)
			//  u=runif(1)<exp(f1-f0)
			//  th0[j]=th1[j]*(u==1)+th0[j]*(u==0)
			//  f0=f1*(u==1)+f0*(u==0)

			if rand.Float64() < exp(f1-f0) {
				th0[j] = th1[j]
				f0 = f1
				arate[j] += 1
			}
			vth[i][j] = th0[j]
		}
	}

	for i, _ := range arate {
		arate[i] /= float64(m)
	}
	return
}
