// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package bayes

// Log posterior of difference and sum of logits in a 2x2 table.

// LogCTablePost returns the log posterior density for the difference and sum of logits in a 2x2 contingency table
// for independent binomial samples and uniform prior placed on the logits.
func LogCTablePost(s1, f1, s2, f2, theta1, theta2 float64) float64 {
	// Arguments
	// theta1 - difference of logits
	// theta2 - sum of logits
	// s1,f1,s2,f2 -  number of successes and failures for first sample, and then the second sample
	// Returns:
	// value of the log posterior

	logitp1 := (theta1 + theta2) / 2
	logitp2 := (theta2 - theta1) / 2
	term1 := s1*logitp1 - (s1+f1)*log(1+exp(logitp1))
	term2 := s2*logitp2 - (s2+f2)*log(1+exp(logitp2))
	return term1 + term2
}

/*&
logctablepost=function (theta, data) 
{
    theta1 = theta[1]
    theta2 = theta[2]
    s1 = data[1]
    f1 = data[2]
    s2 = data[3]
    f2 = data[4]
    logitp1 = (theta1 + theta2)/2
    logitp2 = (theta2 - theta1)/2
    term1 = s1 * logitp1 - (s1 + f1) * log(1 + exp(logitp1))
    term2 = s2 * logitp2 - (s2 + f2) * log(1 + exp(logitp2))
    return(term1 + term2)
}
*/
