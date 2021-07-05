package leetcode

// 计数排序的思路
// 对整个数组遍历了两遍
// 时间复杂度: O(n)
// 空间复杂度: O(k), k为元素的取值范围
func sortColors_1(nums []int) {
	count := map[int]int{} // 存放0, 1, 2三个元素的频率
	for _, num := range nums {
		if _, ok := count[num]; ok {
			count[num]++
		} else {
			count[num] = 1
		}
	}

	index := 0
	for i := 0; i < count[0]; i++ {
		nums[index] = 0
		index++
	}
	for i := 0; i < count[1]; i++ {
		nums[index] = 1
		index++
	}
	for i := 0; i < count[2]; i++ {
		nums[index] = 2
		index++
	}
}
