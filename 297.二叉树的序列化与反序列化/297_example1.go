package LeetCode297

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// bfs
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		if queue[0] != nil {
			res = append(res, strconv.Itoa(queue[0].Val))
			queue = append(queue, queue[0].Left, queue[0].Right) // 左右子树入队列，之后下一层就会遍历
		} else {
			res = append(res, "#")
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	var res = strings.Split(data, ",")
	var root = &TreeNode{} // 根节点的初始化
	root.Val, _ = strconv.Atoi(res[0])
	var queue = []*TreeNode{root}
	res = res[1:]
	for len(queue) > 0 {
		// 每次取0，1 两个值
		if res[0] != "#" { // 表示该节点值不为空
			l, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{Val: l}
			queue = append(queue, queue[0].Left) // 再把左节点入队列，下一层会遍历该左子树
		}

		if res[1] != "#" {
			r, _ := strconv.Atoi(res[1])
			queue[1].Right = &TreeNode{Val: r}
			queue = append(queue, queue[1].Right)
		}
		queue = queue[1:] // 当前节点分析完，出队列
		res = res[2:]
	}
	return root
}
