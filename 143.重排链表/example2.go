package LeetCode143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0)

	//append to array
	curPtr := head
	nextPtr := head
	for curPtr != nil {
		nextPtr = curPtr.Next
		arr = append(arr, curPtr)
		curPtr.Next = nil
		curPtr = nextPtr
	}

	//loop
	lastPtr := head
	maxIdx := len(arr) - 1
	midIdx := maxIdx / 2
	for i := 0; i < midIdx; i++ {
		nextIdx := i + 1
		curPtr = arr[i]
		lastPtr = arr[maxIdx-i]

		curPtr.Next = lastPtr
		lastPtr.Next = arr[nextIdx]
	}

	//last node
	if len(arr)%2 == 0 && len(arr) > 1 {
		arr[midIdx].Next = arr[midIdx+1]
	}
}
