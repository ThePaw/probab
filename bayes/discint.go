// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Highest probability interval for a discrete probability distribution.
// Ref.: Albert (2009): 184 [mnormt.onesided()]

import (
	"sort"
)

// cumSum returns cumulative sums of a slice.
func cumSum(x []float64) []float64 {
	v := make([]float64, len(x))
	for i, _ := range x {
		if i == 0 {
			v[i] = x[i]
		} else {
			v[i] = v[i-1] + x[i]
		}
	}
	return v
}

type IndexSorter struct {
	Target  []float64
	Indices []int
}

func NewSorter(t []float64) IndexSorter {
	iv := make([]int, len(t))
	for i := range iv {
		iv[i] = i
	}
	return IndexSorter{Target: t, Indices: iv}
}
func (s IndexSorter) Len() int           { return len(s.Target) }
func (s IndexSorter) Less(i, j int) bool { return s.Target[i] < s.Target[j] }
func (s IndexSorter) Swap(i, j int) {
	s.Target[i], s.Target[j] = s.Target[j], s.Target[i]
	s.Indices[i], s.Indices[j] = s.Indices[j], s.Indices[i]
}

// DiscHPI computes a highest probability interval for a discrete distribution.  
func DiscHPI(x, p []float64, probContent float64) (probExact float64, hpiSet []float64) {
	// Arguments:
	// x - values where probability is listed
	// p - probability at x
	// probContent - target probability content of the HPI
	// Returns:
	// probExact - exact probability content of the HPI
	// hpiSet set of values of x within the highest probability interval
	s := NewSorter(p)
	sort.Sort(s)
	ix := s.Indices
	ps := s.Target

	// reverse sorted indices
	iRev := make([]int, len(ix))
	for i, _ := range ix {
		iRev[i] = ix[len(ix)-i-1]
	}

	// reverse sorted probabilities
	pRev := make([]float64, len(ps))
	for i, _ := range ps {
		pRev[i] = ps[len(ps)-i-1]
	}

	// sort x
	xRev := make([]float64, len(iRev))
	for i, _ := range xRev {
		xRev[i] = x[iRev[i]]
	}

	cp := cumSum(pRev)

	// find first index where cp>=probContent
	j := 0
	for i, _ := range cp {
		if cp[i] >= probContent {
			break
		}
		j++
	}
	probExact = cp[j]
	hpiSet = make([]float64, j+1)
	for i := 0; i < j+1; i++ {
		hpiSet[i] = xRev[i]
	}
	sort.Float64s(hpiSet)
	return probExact, hpiSet
}
