package removenelem

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// via array
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := make([]*ListNode, 0) // 30 -- max size by task

	// transform linked list to arrat of pointers
	node := head
	for node != nil {
		list = append(list, node)
		node = node.Next
	}
	if len(list) == 1 {
		// corner case
		return nil
	}
	if n == len(list) {
		// remove from start
		return head.Next
	}
	prev := list[len(list)-n-1]
	var next *ListNode
	if n != 1 {
		// remove not from end
		next = list[len(list)-n+1]
	}

	prev.Next = next

	return head
}
