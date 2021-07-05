package leetcode

// 滑动窗口的思路
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func minSubArrayLen_3(target int, nums []int) int {
	l, r := 0, -1 // nums[l...r]为我们的滑动窗口
	sum := 0
	res := len(nums) + 1
	for ; l < len(nums); { // 窗口的左边界在数组范围内,则循环继续
		if r < len(nums)-1 && sum < target {
			r++
			sum = sum + nums[r]
		} else { // r已经到头 或者 sum >= s
			sum = sum - nums[l]
			l++
		}
		if ln := r - l + 1; sum >= target && ln < res {
			res = ln
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
