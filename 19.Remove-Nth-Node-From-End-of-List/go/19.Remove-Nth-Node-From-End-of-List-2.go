package leetcode

// 使用双指针, 对链表只遍历了一遍
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeNthFromEnd_2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	point1, point2 := dummyHead, dummyHead
	for i := 0; i < n + 1; i++ {
		point2 = point2.Next
	}

	for ; point2 != nil; {
		point1 = point1.Next
		point2 = point2.Next
	}


	delNode := point1.Next
	point1.Next = delNode.Next
	delNode.Next = nil

	retNode := dummyHead.Next
	dummyHead.Next = nil

	return retNode
}
