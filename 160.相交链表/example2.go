package LeetCode160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapInterface := make(map[interface{}]interface{})
	for headA != nil {
		mapInterface[headA] = 1
		headA = headA.Next
	}

	for headB != nil {
		if mapInterface[headB] == 1 {
			return headB
		}

		headB = headB.Next
	}

	return nil
}
