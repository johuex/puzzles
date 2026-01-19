package deleteduplicates

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var prevPtr *ListNode
	actPtr := head

	for actPtr != nil {
		if prevPtr != nil && prevPtr.Val == actPtr.Val {
			prevPtr.Next = actPtr.Next
			actPtr = actPtr.Next
		} else {
			prevPtr = actPtr
			actPtr = actPtr.Next
		}

	}
	return head
}
