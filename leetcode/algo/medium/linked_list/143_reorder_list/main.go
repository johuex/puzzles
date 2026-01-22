package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	node := head
	list := make([]*ListNode, 0)

	// make additional array for fast random access
	for node != nil {
		list = append(list, node)
		node = node.Next
	}

	arrayPtr := len(list) - 1
	node = head
	var idx int
	// make reordering
	for node != list[arrayPtr] && node.Next != list[arrayPtr] {
		oldNext := node.Next        // remember
		node.Next = list[arrayPtr]  // insert new item
		node.Next.Next = oldNext    // conect inserted with rmembered
		list[arrayPtr-1].Next = nil // cut elem in linked list
		arrayPtr--
		idx++
		node = oldNext // jump over inserted node
	}
}
