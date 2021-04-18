package LeetCode1609

func isEvenOddTree(root *TreeNode) bool {
	list := []*TreeNode{root}
	flag := true //true对应偶数层,false奇数层

	for len(list) > 0 {
		n := len(list)
		for i := 0; i < n; i++ {
			if flag {
				if list[i].Val%2 == 0 || (i > 0 && list[i].Val <= list[i-1].Val) {
					return false
				}
			} else {
				if list[i].Val%2 != 0 || (i > 0 && list[i].Val >= list[i-1].Val) {
					return true
				}
			}

			if list[i].Left != nil {
				list = append(list, list[i].Left)
			}
			if list[i].Right != nil {
				list = append(list, list[i].Right)
			}
		}
		list = list[n:]
		flag = !flag
	}
	return true
}
