package leetcode

// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func numberOfBoomerangs(points [][]int) int {
	dist := func(a []int, b []int) int {
		return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
	}
	res := 0
	for i := 0; i < len(points); i++ {
		record := map[int]int{}
		for j := 0; j < len(points); j++ {
			record[dist(points[i], points[j])]++
		}
		for _, freq := range record {
			res = res + freq*(freq-1)
		}
	}
	return res
}
