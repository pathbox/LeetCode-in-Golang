package main

import "fmt"

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5
	// for head != nil {
	// 	fmt.Printf("%d -> ", head.Val)
	// 	head = head.Next
	// }

	pre := reversrList(head)

	for pre != nil {
		fmt.Printf("%d -> ", pre.Val)
		pre = pre.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表的实现
func reversrList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	p2 := head.Next
	var p3 *ListNode = nil

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1 // 反向指针，后移动节点
		p1 = p2
		p2 = p3
	}

	head.Next = nil
	head = p1

	return head
}

func printNode(head *ListNode) {
	for head != nil {
		//fmt.Print(head.value, "\t")
		fmt.Println(head)
		head = head.Next
	}
	fmt.Println()
}

/*
pre是cur的最前面那位（pre = cur）
cur就是当前位的后面链表元素（cur = cur.Next）
cur.Next肯定是接pre（cur.Next = pre）
*/
