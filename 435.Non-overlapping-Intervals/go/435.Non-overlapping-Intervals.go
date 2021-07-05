package leetcode

import "sort"

// 贪心算法
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}
