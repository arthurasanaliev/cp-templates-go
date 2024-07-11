func sieve(n int) []int {
  isPrime := make([]bool, n + 1)
  for i := range isPrime {
    isPrime[i] = true
  }
  isPrime[0], isPrime[1] = false, false
  for i := 2; i * i <= n; i++ {
    if isPrime[i] {
      for j = i * i; j <= n; j += i {
        isPrime[j] = false
      }
    }
  }
  var primes []int
  for i := 2; i <= n; i++ {
    if isPrime[i] {
      primes = append(primes, i)
    }
  }
  return primes
}
