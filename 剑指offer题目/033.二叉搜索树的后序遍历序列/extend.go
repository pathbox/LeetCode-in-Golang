// https://www.bilibili.com/video/BV15D4y1X7Tt?from=search&seid=9781582127527481898
package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

func main() {
	posArr := []int{2, 4, 3, 6, 8, 7, 5}
	r := process(posArr, 0, len(posArr)-1)
	fmt.Println(r)
}

// 已知一个搜索二叉树后续遍历的数组posArr，重建出整个树返回head节点
func process(posArr []int, L, R int) *Node {
	if L > R {
		return nil
	}
	head := &Node{Val: posArr[R]} // 得到当前层子树的head节点，去得到head的left和right子树
	M := L - 1
	left := L
	right := R - 1

	for left <= right {
		mid := left + (right-left)>>1 // 这种方式是避免溢出的最好方式
		// 目的是找到大于posArr[R]的最小分界位置
		if posArr[mid] < posArr[R] {
			M = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	head.Left = process(posArr, L, M)
	head.Right = process(posArr, M+1, R-1)
	return head
}

//O(nlogn)
