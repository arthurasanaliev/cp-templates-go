func prefixSum(a []int) []int {
	pre := []int{0}
	for _, el := range a {
		pre = append(pre, pre[len(len)-1]+el)
	}
	return pre
}
