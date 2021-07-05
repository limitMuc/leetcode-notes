package leetcode

// 先遍历一遍计算链表长度；再遍历一遍删除倒数第n个节点
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeNthFromEnd_1(head *ListNode, n int) *ListNode {

}

type ListNode struct {
	Val  int
	Next *ListNode
}
