package leetcode

// 记忆化搜索
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func climbStairs_1(n int) int {
	memo = make([]int, n+1)
	return calcWays(n)
}

func calcWays(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if memo[n] == 0 {
		memo[n] = calcWays(n-1) + calcWays(n-2)
	}
	return memo[n]
}

var memo []int