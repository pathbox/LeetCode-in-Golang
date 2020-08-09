package LeetCode1361

/*
1、如果leftChild和rightChild里出现重复元素了，返回错误；
2、找到入度为空的元素装进队列q，如果q的长度不等于1，说明有多个树根或没有树根，返回错误；
3、从树根开始宽度优先遍历，逐个把遍历到的节点加入队列q，最后返回的时候验证是否遍历到了所有节点，排除多棵树的情况
*/
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	v := make([]bool, n)
	for i := 0; i < n; i++ {
		for _, j := range []int{leftChild[i], rightChild[i]} {
			if j != -1 {
				if v[j] {
					return false
				}
				v[j] = true
			}
		}
	}
	q := []int{}
	for i, j := range v {
		if !j {
			q = append(q, i)
		}
	}
	if len(q) != 1 {
		return false
	}
	for i := 0; i < len(q); i++ {
		if leftChild[q[i]] != -1 {
			q = append(q, leftChild[q[i]])
		}
		if rightChild[q[i]] != -1 {
			q = append(q, rightChild[q[i]])
		}
	}
	return len(q) == n
}
