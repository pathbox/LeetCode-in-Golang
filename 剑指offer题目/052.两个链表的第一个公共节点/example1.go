package Offer052

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//创建一个临时的map，map里面存的是地址，外面存的是值
	var tmpMap = make(map[*ListNode]int)
	//如果AB都空，则跳出循环
	for headA != nil || headB != nil { // 一个for循环中解决

		if headA != nil {
			//如果map中没有该节点，则把该节点添加进去
			if _, ok := tmpMap[headA]; !ok {
				//令当前节点的值为headA.Val
				tmpMap[headA] = headA.Val
			} else {
				return headA
			}
			//指向下一个
			headA = headA.Next
		}

		if headB != nil {
			//如果map中没有该节点，则把该节点添加进去
			if _, ok := tmpMap[headB]; !ok {
				//令当前节点的值为headB.Val
				tmpMap[headB] = headB.Val
			} else {
				return headB
			}

			headB = headB.Next
		}
	}

	return nil
}
