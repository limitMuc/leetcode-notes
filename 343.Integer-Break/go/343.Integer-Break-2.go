package leetcode

// 动态规划
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func integerBreak_2(n int) int {
	memo := make([]int, n+1)
	memo[1] = 1
	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			memo[i] = max3(memo[i], j*(i-j), j*memo[i-j])
		}
	}

	return memo[n]
}
