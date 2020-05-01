package LeetCode173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root: root}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	i, _ := this.n(true)
	return i
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	_, ok := this.n(false)
	return ok
}

func (this *BSTIterator) n(move bool) (int, bool) {
	var max *TreeNode

	for this.root != nil {
		if this.root.Left == nil {
			if move {
				defer func() {
					this.root = this.root.Right
				}()
			}
			return this.root.Val, true

		} else {
			//寻找左树最大节点
			max = this.root.Left
			for max.Right != nil {
				max = max.Right
			}

			//最大节点指向root
			max.Right = this.root

			//root = root.Left :移动root到root.left
			//root.Left = nil  :砍左子树，避免下一次遍历到root时，再进入到左子树
			this.root, this.root.Left = this.root.Left, nil
		}
	}
	return 0, false
}
