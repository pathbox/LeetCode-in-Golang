package LeetCode1361

// 根据图论，要满足树的条件，需要满足edge = vertices - 1edge=vertices−1
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	edge := 0
	for _, value := range leftChild {
		if value != -1 {
			edge++
		}
	}

	for _, value := range rightChild {
		if value != -1 {
			edge++
		}
	}

	if edge == n-1 {
		return true
	} else {
		return false
	}
}
