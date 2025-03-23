package introduction

import "math"

func isPrime(num int) bool {

	if num == 1 || num == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))

	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}
