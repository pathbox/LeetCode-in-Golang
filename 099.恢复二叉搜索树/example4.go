package LeetCode099

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
https://leetcode.wang/leetCode-94-Binary-Tree-Inorder-Traversal.html
根节点初始化为current

如果current的左节点为空，输出该节点，同时current = current.Right

如果current的左节点不为空，那么先去找到这个节点左子树的最右边节点leftRight

首先leftRight = current.Left如果leftRight的右节点不为空且不为current，则循环leftRight = leftRight.Right
如果leftRight的右节点为空（也就是没有指向current），则leftRight.Right = current current = current.Left（leftRight右节点指向当前current，current进入左子树）
如果pre的右节点为current，意味着这个current的左子树是完全遍历过了,那么恢复这个右节点leftRight.Right = nil，输出当前节点，current进入右子树，current = current.right
重复2、3直到current为nil

*/

//方法二：Morris方法中序遍历+冒泡
func recoverTree(root *TreeNode) {
	current := root
	var prevent *TreeNode
	var firstNode *TreeNode
	var secondNode *TreeNode
	for current != nil {
		//如果左子树为空那么说明这个节点是某个子树最左边的节点，可以遍历这个节点
		if current.Left == nil {
			//fmt.Println(prevent,current.Val)
			if prevent != nil && prevent.Val > current.Val {
				if firstNode == nil {
					firstNode = prevent
				}
				secondNode = current
			}
			prevent = current
			//同时进入它的右子树
			current = current.Right
		} else {
			//不为空说明还可以向左进行，首先找到current左子树最右边的节点leftRight的右子树接到current上
			leftRight := current.Left
			for leftRight.Right != nil && leftRight.Right != current {
				leftRight = leftRight.Right
			}
			//如果左子树没有被遍历过,就把pre的右节点指向current,并把current向左迈一步
			if leftRight.Right != current {
				leftRight.Right = current
				current = current.Left
			} else {
				//如果遍历过，就把这个pre的指向去掉，恢复原来树结构，并让current进入右子树
				leftRight.Right = nil
				//fmt.Println(prevent,current.Val)
				if prevent != nil && prevent.Val > current.Val {
					if firstNode == nil {
						firstNode = prevent
					}
					secondNode = current
				}
				prevent = current
				current = current.Right
			}
		}
	}
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}
