package leetcode

// 动态规划
// 时间复杂度: O(len(nums) * O(sum(nums)))
// 空间复杂度: O(len(nums) * O(sum(nums)))
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}

	if sum%2 != 0 {
		return false
	}

	ln := len(nums)
	C := sum / 2
	memo := make([]bool, C+1)

	for i := 0; i <= C; i++ {
		memo[i] = (nums[0] == i)
	}

	for i := 1; i < ln; i++ {
		for j := C; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j-nums[i]]
		}
	}

	return memo[C]
}
