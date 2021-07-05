package leetcode

// 时间复杂度: O(n^2)
// 空间复杂度: O(n^2)
func fourSumCount_1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	record := map[int]int{}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			record[n3+n4]++
		}
	}

	res := 0
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			res = res + record[-(n1+n2)]
		}
	}
	return res
}
