// Copyright 2012 The Probab Authors. All rights reserved. See the LICENSE file.

package dst

// Gaussian ratio distribution. 

//GearyHinkleyTransformation transforms the ratio of two normally distributed variables to the transformed variable T would approximately have a standard Gaussian distribution. See Hinkley(1969). 
func GearyHinkleyTransformation(μX, σX, μY, σY float64) float64 {
	//A Geary–Hinkley transformation, under certain assumptions, returns the transformed variable T that would approximately have a standard Gaussian distribution. The approximation is good if Y is unlikely to assume negative values.
	// X = N(μX, σ2X) and Y = N(μY, σ2Y) 
	// Z = X/Y
	// http://en.wikipedia.org/wiki/Ratio_distribution#A_transformation_to_Gaussianity

	t1 := muY*z - muX
	t2 := math.Sqrt(σ2Y*z*z - 2*ρ*σX*σY*z + σ2X)
	return t1 / t2
}
