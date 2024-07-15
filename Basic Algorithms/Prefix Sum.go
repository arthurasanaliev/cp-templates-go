func prefixSum(a []int) []int {
	pre := make([]int, len(a)+1)
	for i, el := range a {
		pre[i+1] = pre[i] + el
	}
	return pre
}
