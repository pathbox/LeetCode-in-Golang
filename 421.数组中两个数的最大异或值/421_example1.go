package LeetCode421

import "strings"

type trieNode struct {
	children []*trieNode
}

func constructNode() *trieNode {
	return &trieNode{children: make([]*trieNode, 2)}
}

// 将整数转为32位的二进制字符串，前面不为1用0补充
func toBinaryString(num int) string {
	if num == 0 {
		return "0"
	}
	bins := make([]rune, 32)
	mask := 1 << 31
	for i := 0; i < len(bins); i++ {
		if (mask & num) != 0 {
			bins[i] = '1'
		} else {
			bins[i] = '0'
		}
		mask >>= 1 // 向右移动一位，去掉一个最高位
	}
	return strings.TrimLeft(string(bins), "0") // 左边空白字符串用0补充
}

func findMaximumXOR(nums []int) int {
	maxNum := 0
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
	}
	l := len(toBinaryString(maxNum))
	mask := 1 << l
	bins := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		bins[i] = toBinaryString(nums[i] | mask)[1:]
	}
	maxXor := 0
	root := constructNode()
	for _, bin := range bins {
		node, xorNode := root, root
		xor := 0
		for _, r := range bin {
			// 插入
			idx := r - '0'
			if node.children[idx] == nil {
				node.children[idx] = constructNode()
			}
			node = node.children[idx]

			// 计算xor
			if xorNode.children[1-idx] != nil {
				xor = (xor << 1) + 1
				xorNode = xorNode.children[1-idx]
			} else {
				xor = xor << 1
				xorNode = xorNode.children[idx]
			}
		}
		if maxXor < xor {
			maxXor = xor
		}
	}
	return maxXor
}
