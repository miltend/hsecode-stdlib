package math

//package main

import "math"

func NthPrime(n int) int {
	switch {
	case n <= 0:
		panic("Panicking")
	case n <= 141:
		var firstPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
			73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179,
			181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283,
			293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419,
			421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547,
			557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661,
			673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811}
		return firstPrimes[n-1]
	case n < 24368:
		var maxSize = int(math.Round(math.Pow(float64(n), 1.36)))
		isPrime := make([]bool, maxSize)
		var primes []int
		for i := 2; i < maxSize; i++ {
			if isPrime[i] == true {
				continue
			}
			primes = append(primes, i)
			for k := i * i; k < maxSize; k += i {
				isPrime[k] = true
			}

		}
		return primes[n-1]
	default:
		const maxSize = 200000000
		isPrime := make([]bool, maxSize)
		var primes []int
		for i := 2; i < maxSize; i++ {
			if isPrime[i] == true {
				continue
			}
			primes = append(primes, i)
			for k := i * i; k < maxSize; k += i {
				isPrime[k] = true
			}
		}
		return primes[n-1]
	}

	return 0
}
