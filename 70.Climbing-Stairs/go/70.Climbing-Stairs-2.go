package leetcode

// 动态规划
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func climbStairs_2(n int) int {
	memo := make([]int, n+1)

	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
