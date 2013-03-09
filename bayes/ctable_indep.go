// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LyRowsCENSE file.

package bayes

// compute Bayes factor against independence
// using Albert and Gupta independence priors
// y - matrix of counts
// k - Dirichlet precision hyperparameter
// m - number of simulations
// Ref.: Albert (2009)

import (
	"code.google.com/p/probab/dst"
	"fmt"
)

// rdirichlet simulates a sample from a Dirichlet distribution.
func rdirichlet(m int, par []float64) [][]float64 { //// should accept vector, not mtx!! rewrite! Pass mtx unloaded colwise to vec
	// m number of simulations required
	// par vector of parameters of the Dirichlet distribution

	// make outMatrix
	nCol := len(par)
	s := make([]float64, m*nCol)
	out := make([][]float64, m)
	for i, p := 0, 0; i < m; i++ {
		out[i] = s[p : p+nCol]
		p += nCol
	}

	for i := range out { // nRows
		dir := dst.DirichletNext(par)
		for j, _ := range dir {
			out[i][j] = dir[j]
		}
	}
	return out
}

// ldirichlet returns the log dirichlet function for multiple values stored in rows of matrix a.
func ldirichlet(a [][]float64) []float64 {
	n := len(a)
	v := make([]float64, n)
	for i := 0; i < n; i++ {
		v[i] = lDirich(a[i])
	}
	return v
}

// FactCTableIndep returns a Bayes factor against independence for a two-way contingency table assuming a 
// "close to independence" alternative model.
func FactCTableIndep(y [][]float64, k float64, m int) (bf, nse float64) {
	// Arguments
	// y matrix of counts
	// k Dirichlet precision hyperparameter
	// m number of simulations
	//
	// Return values
	// bf value of the Bayes factor against hypothesis of independence
	// nse estimate of the simulation standard error of the computed Bayes factor

	yc := colSums(y)
	yr := rowSums(y)
	n := 0.0
	for _, val := range yc {
		n += val
	}

	//d=dim(y); 
	yRows := len(y)
	yCols := len(y[0])

	yr1 := make([]float64, len(yr))
	for i, _ := range yr1 {
		yr1[i] = yr[i] + 1
		//fmt.Println("yr1[i]: ", yr1[i])
	}

	yc1 := make([]float64, len(yc))
	for i, _ := range yc1 {
		yc1[i] = yc[i] + 1
		//fmt.Println("yc1[i]: ", yc1[i])
	}

	etaA := rdirichlet(m, yr1)
	etaB := rdirichlet(m, yc1)

	// make keta
	nCol := yRows * yCols
	s1 := make([]float64, m*nCol)
	keta := make([][]float64, m)
	for i, p := 0, 0; i < m; i++ {
		keta[i] = s1[p : p+nCol]
		p += nCol
	}

	// make ketaY
	s2 := make([]float64, m*nCol)
	ketaY := make([][]float64, m)
	for i, p := 0, 0; i < m; i++ {
		ketaY[i] = s2[p : p+nCol]
		p += nCol
	}
	fmt.Println("Dim keta: ", len(keta), len(keta[0]))

	// fill in keta, ketaY
	//keta=cbind(keta,k*etaA[,i]*etaB[,j])
	col := 0
	for i, _ := range y {
		for j, _ := range y[0] {
			for h := range keta { // nRows
				keta[h][col] = k * etaA[h][i] * etaB[h][j]
				ketaY[h][col] = k*etaA[h][i]*etaB[h][j] + y[i][j]
			}
			col++
		}
	}

	logint := ldirichlet(ketaY)
	ld2 := ldirichlet(keta)

	for i, _ := range keta { //m
		logint[i] -= ld2[i]
	}

	for i, _ := range yr {
		for j, _ := range logint { //m
			logint[j] -= yr[i] * log(etaA[j][i])
		}
	}

	for i, _ := range yc {
		for j, _ := range logint { //m
			logint[j] -= yc[i] * log(etaB[j][i])
		}
	}

	zz := make([]float64, len(logint))
	for i, _ := range logint {
		zz[i] = exp(logint[i])
	}

	bf, nse = meanSd(zz)
	nse /= sqrt(float64(m))
	return
}
