package LeetCode1361

/*
计算每个节点的入度
除根节点外每个节点入度均为1
注意：孤立节点属于二叉树
*/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	inDegree := make(map[int]int, 0)
	inDegree[0] = 0
	for i := 0; i < len(leftChild); i++ {
		if leftChild[i] != -1 {
			inDegree[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			inDegree[rightChild[i]]++
		}
	}

	// 单个节点
	if len(inDegree) == 1 && inDegree[0] == 1 {
		return true
	}

	if inDegree[0] != 0 {
		return false
	}

	for i := 1; i < len(leftChild); i++ {
		if inDegree[i] != 1 { // 除根节点外每个节点入度均为1
			return false
		}
	}
	return true
}
