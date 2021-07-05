package leetcode

// 对撞指针
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func twoSum_3(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for ; l < r; {
		res := numbers[l] + numbers[r]
		if res == target {
			return []int{l + 1, r + 1}
		} else if res > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}
