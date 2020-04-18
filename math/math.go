package math

func NthPrime(n int) int {
	if n <= 0 {
		panic("Panicking")
	}
	const maxSize = 1000005
	primes := map[int]int{}
	isPrime := map[int]bool{}
	for i := 1; i < maxSize; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p < maxSize; p++ {
		if isPrime[p] == true {
			for k := p * p; k < maxSize; k += p {
				isPrime[k] = false
			}
		}
	}
	i := 1
	for p := 2; p < maxSize; p++ {
		if isPrime[p] == true {
			primes[i] = p
			i++
		}
	}

	return primes[n]
}
