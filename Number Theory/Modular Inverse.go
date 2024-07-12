func binPow(a, b, m int) int {
  res := 1
  for b > 0 {
    if b % 2 == 1 {
      res = res * a % m;
    }
    a = a * a % m;
    b /= 2
  }
  return res
}

func inverse(a, m int) int {
  res := binPow(a, m - 2, m)
  return res
}
