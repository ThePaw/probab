// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Density function of a histogram distribution.
// Ref.: Albert (2009)

// HistPrior returns the density of a probability distribution defined on a set of equal-width intervals.
func HistPrior(p, midpts, prob []float64) []float64 {
	// Arguments:
	// p - vector of values for which density is to be computed
	// midpts - vector of midpoints of the intervals
	// prob - vector of probabilities of the intervals
	// Returns:
	// vector of values of the probability density

	maxDiff := 1e4
	if len(midpts) != len(prob) {
		panic(" len(midpts) != len(prob)")
	}

	for i, _ := range midpts {
		d := midpts[1] - midpts[0]
		if i > 0 {
			if abs(midpts[i]-midpts[i-1]-d) > maxDiff {
				panic("midpts not equidistant")
			}
		}

	}

	lo := make([]float64, len(midpts))
	val := make([]float64, len(p))
	binwidth := midpts[1] - midpts[0]
	for i, _ := range midpts {
		lo[i] = midpts[i] - binwidth/2
	}

	for i, _ := range p {
		for j := 0; j < len(lo); j++ {
			if p[i] >= lo[j] {
				val[i] = prob[j]
			}
		}
	}
	return val
}
