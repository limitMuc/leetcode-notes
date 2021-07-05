package leetcode

// 原地(in place)解决该问题
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func moveZeroes_3(nums []int)  {
	k := 0// nums中, [0...k)的元素均为非0元素

	// 遍历到第i个元素后,保证[0...idx]中所有非0元素
	// 都按照顺序排列在[0...k)中
	// 同时, [k...idx] 为 0
	for idx, num := range nums {
		if num != 0 {
			nums[k], nums[idx] = nums[idx], nums[k]
			k++
		}
	}
}