// Copyright 2012 - 2013 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// PropDisc returns the posterior distribution for a proportion for a discrete prior distribution.
func PropDisc(p, prior []float64, succ, fail int) []float64 {
	//Arguments: 
	// p vector of proportion values
	// prior vector of prior probabilities
	// succ number of successes
	// fail number of failures
	// Value:
	// vector of posterior probabilities.
	// Ref.: Albert (2009): Chapter 2.3: 19-22.

	s := float64(succ)
	f := float64(fail)
	mx := -1e99

	p1 := make([]float64, len(p))
	for i, _ := range p {
		if p[i] == 0 || p[i] == 1 {
			p1[i] = 0.5
		} else {
			p1[i] = p[i]
		}
	}

	like := make([]float64, len(p))
	for i, _ := range like {
		like[i] = s*log(p1[i]) + f*log(1-p1[i])
	}

	for i, _ := range like {
		if !(p[i] > 0 && p[i] < 1) {
			if (p[i] == 0 && s > 0) || (p[i] == 1 && f > 0) {
				like[i] = nan
			}
		}
	}

	for i, _ := range like {
		if like[i] > mx {
			mx = like[i]
		}
	}

	for i, val := range like {
		like[i] = exp(val - mx)
	}

	//product=like*prior
	product := make([]float64, len(p))
	for i, _ := range product {
		product[i] = like[i] * prior[i]
	}

	//post=product/sum(product)
	sum := 0.0
	for i, _ := range product {
		sum += product[i]
	}

	post := make([]float64, len(product))
	for i, _ := range product {
		post[i] = product[i] / sum
	}
	return post
}
