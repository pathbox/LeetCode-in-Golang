package LeetCode449

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

//bfs
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
			queue = append(queue, queue[0].Left, queue[0].Right)
		} else {
			res = append(res, "#") // 表示nil节点
		}
		queue = queue[1:]
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializes(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	var res = strings.Split(data, ",")
	var root = &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	var queue = []*TreeNode{root}
	res = res[1:]
	// 每次从res取两个数，反序列化为左右子树节点
	for 0 < len(queue) {
		if res[0] != "#" {
			l, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{Val: l}
			queue = append(queue, queue[0].Left) // 入queue，下一次遍历
		}

		if res[1] != "#" {
			r, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNode{Val: r}
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		res = res[2:]
	}
	return root
}
