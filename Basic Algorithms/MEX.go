import "sort"

func MEX(a []int) int {
  sort.Ints(a)
  if a[0] != 0 {
    return 0
  }
  for i := 1; i < len(a); i++ {
    if a[i] > a[i - 1] + 1 {
      return a[i - 1] + 1
    }
  }
  return a[len(a) - 1] + 1
}
