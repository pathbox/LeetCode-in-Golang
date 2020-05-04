package LeetCode654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 依次插入，先选取首项为root
接下来每一项：如果小于root，则插入root的右树（insertRight），如果大于root，则选取该点为新root，将原树变为新root的左树（reRoot）
实现insertRight与reRoot函数
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := newNode(nums[0])

	for _, val := range nums[1:] {
		root = insert(val, root)
	}

	return root
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func insert(val int, root *TreeNode) *TreeNode {
	// 插入右子树
	if val < root.Val {
		insertRight(val, root)
		return root
	}

	return reRoot(val, root)

}

// 对于路径上的节点，如果遇到小于该值的节点，则用该值取代原节点，并将该节点的子树赋值给新节点的左树
func insertRight(val int, root *TreeNode) {
	cursor := root
	for {
		if cursor.Right != nil {
			if cursor.Right.Val < val {
				cursor.Right = reRoot(val, cursor.Right)
				return
			}
			cursor = cursor.Right // 不断的去找到合适的右子树的位置
		} else {
			break
		}
	}

	cursor.Right = newNode(val)
	return
}

// 修改根为新节点并将原树赋值为新root的左树
func reRoot(val int, root *TreeNode) (newRoot *TreeNode) {
	newRoot = newNode(val)
	newRoot.Left = root
	return
}
