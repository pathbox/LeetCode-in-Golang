import "fmt"
package main


//将BST转化为双向循环链表，不允许新建节点
//为防止歧义，左指针表示双链表向前指，右指针表示双链表向后指
type TreeNode struct {
    Val   int
    Left * TreeNode
    Right * TreeNode
}

var pre *TreeNode // 必须在全局变量上才可以实现
// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/solution/mian-shi-ti-36-er-cha-sou-suo-shu-yu-shuang-xian-5/
func treeToDoublyList(root * TreeNode) * TreeNode {
    if root == nil {
        return nil
    }
    helper(root)

    // 处理头尾节点
    head, tail: = root, root
    for head.Left != nil {
        head = head.Left // 最左节点是头节点
    }
    for head.Right != nil {
        tail = tail.Right // 最右节点是尾节点
    }
    head.Left = tail
    tail.Right = head.Right
    return head
}

func helper(root *TreeNode) {
    if root == nil {
        return
    }
    // 中序遍历：递增
    helper(root.Left) // 会递归到最后一个左子树，左节点，然后开始返回处理
    if pre != nil {
        root.Left = pre
        pre.Right = root
    }
    pre = root
    helper(root.Right)
}

func main() {
	root := &TreeNode{4, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{5, nil, nil}
	node3 := &TreeNode{1, nil, nil}
	node4 := &TreeNode{3, nil, nil}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	head := treeToDoublyList(root)
	tail := head.Left
	//从头开始遍历
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", head.Val)
		head = head.Right
	}
	//从尾开始遍历
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d\t", tail.Val)
		tail = tail.Left
	}

}
