package leetcode

// 时间复杂度：0(n)
// 空间复杂度：0(n)
func moveZeroes_1(nums []int) {
	nozero := make([]int, 0, len(nums))
	for _, num := range nums {
		if num != 0 {
			nozero = append(nozero, num) // 将所有非0元素放入nozero
		}
	}
	for i := 0; i < len(nozero); i++ { // 将nozero中的所有元素依次放入到nums开始的位置
		nums[i] = nozero[i]
	}
	for i := len(nozero); i < len(nums); i++ {
		nums[i] = 0 // 将nums剩余的位置放置为0
	}
}
