package removeelements

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// we need to store prev elem and actual to easy erase elem from linked list
	var prevPtr *ListNode
	actPtr := head

	for actPtr != nil {
		if actPtr.Val == val {
			if prevPtr == nil {
				// start of linked list
				head = actPtr.Next
				actPtr = head
				continue
			} else if actPtr.Next == nil {
				// end of linked list
				prevPtr.Next = nil
				actPtr = nil
				continue
			} else {
				// common case, in the middle of linked list
				prevPtr.Next = actPtr.Next
				actPtr = actPtr.Next
			}
		} else {
			prevPtr = actPtr
			actPtr = actPtr.Next
		}

	}
	return head
}
