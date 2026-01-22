package linkedlistcycle

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// first way via map
func detectCycleMap(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ListNode]struct{}) // here we will store pointers to all passed nodes
	m[head] = struct{}{}              // init start of list

	for {
		next := head.Next
		if next == nil {
			// no cycle, return empty node
			return nil
		}
		if _, ok := m[next]; ok {
			// we found elem that looped
			return next
		}
		m[next] = struct{}{}
		head = next
	}
}

func detectCyclePointers(head *ListNode) *ListNode {

	if head == nil {
		// empty list
		return nil
	}

	slowPtr := head
	fastPtr := head

	for {
		if fastPtr.Next == nil || fastPtr.Next.Next == nil {
			// no loop
			return nil
			// there is no point to check slowPtr because it is slower than fastPtr

		}
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next

		if slowPtr == fastPtr {
			// first time pointers meeting
			// set slow to start
			slowPtr = head
			for slowPtr != fastPtr {
				// with identical speed both pointers will meet at start of loop
				slowPtr = slowPtr.Next
				fastPtr = fastPtr.Next
			}
			return fastPtr

		}
	}

}
