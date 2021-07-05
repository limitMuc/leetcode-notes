package leetcode

// 原地(in place)解决该问题
// 时间复杂度：0(n)
// 空间复杂度：0(1)
func moveZeroes_2(nums []int) {
	k := 0 // nums中, [0...k)的元素均为非0元素
	for _, num := range nums { // 遍历到第i个元素后,保证[0...i]中所有非0元素都按照顺序排列在[0...k)中
		if num != 0 {
			nums[k] = num
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0 // 将nums剩余的位置放置为0
	}
}
