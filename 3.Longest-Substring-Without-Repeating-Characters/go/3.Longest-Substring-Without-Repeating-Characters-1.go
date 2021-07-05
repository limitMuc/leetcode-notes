package leetcode

// 滑动窗口
// 时间复杂度: O(len(s))
// 空间复杂度: O(len(charset))
func lengthOfLongestSubstring_1(s string) int {
	l, r := 0, -1 //滑动窗口为s[l...r]
	res := 0
	freq := make([]int, 256, 256)

	// 整个循环从 l == 0; r == -1 这个空窗口开始
	// 到l == s.size(); r == s.size()-1 这个空窗口截止
	// 在每次循环里逐渐改变窗口, 维护freq, 并记录当前窗口中是否找到了一个新的最优值
	for ; l < len(s); {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else { // r已经到头或者子串含有重复字符
			freq[s[l]]--
			l++
		}
		if ln := r - l + 1; ln > res {
			res = ln
		}
	}
	return res
}
