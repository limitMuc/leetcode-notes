package leetcode

// 记忆化搜索
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func integerBreak_1(n int) int {
	memo = make([]int, n+1)
	return breakInteger(n)
}

var memo []int

// 将n进行分割(至少分割两部分), 可以获得的最大乘积
func breakInteger(n int) int {
	if n == 1 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	res := -1
	for i := 1; i <= n-1; i++ {
		res = max3(res, i*(n-i), i*breakInteger(n-i))
		memo[n] = res
	}
	return res
}

func max3(a, b, c int) int {
	max := a
	if a < b {
		max = b
	}
	if max < c {
		max = c
	}
	return max
}
