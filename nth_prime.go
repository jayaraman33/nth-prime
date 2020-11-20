/*package prime

// Nth gives n'th prime number
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	primeCount := 0
	i := 2
	for ; primeCount < n; i++ {
		if isPrime(i) {
			primeCount++
		}
	}
	return i - 1, true
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i <= n/2; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}*/

package prime

// Nth returns the nth prime
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	count := 1
	lastPrime := 2
	for i := 3; count < n; i += 2 {
		prime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			lastPrime = i
			count++
		}
	}
	return lastPrime, true
}
