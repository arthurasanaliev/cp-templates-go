func lastTrue(l, r int) int {
  l--
  for l < r {
    m := l + (r - l + 1) / 2
    if f(m) {
      l = m
    } else {
      r = m - 1;
    }
  }
  return l
}

func firstTrue(l, r int) int {
  r++
  for l < r {
    m := l + (r - l) / 2
    if f(m) {
      r = m
    } else {
      l = m + 1
    }
  }
  return l
}
