package leetcode

import "sort"

// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)
func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	res := 0
	for si, gi := 0, 0; si < len(s) && gi < len(g); {
		if s[si] >= g[gi] {
			res++
			si++
			gi++
		} else {
			gi++
		}
	}
	return res
}
