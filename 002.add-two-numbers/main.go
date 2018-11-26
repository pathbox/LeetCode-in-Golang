package main

import "fmt"

type question struct {
	p para
	a ans
}

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

func main() {
	qs := []question{
		question{
			p: para{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
		question{
			p: para{
				one: makeListNode([]int{9, 8, 7, 6, 5}),
				two: makeListNode([]int{1, 1, 2, 3, 4}),
			},
			a: ans{
				one: makeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		question{
			p: para{
				one: makeListNode([]int{0}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		fmt.Printf("%d = %d\n", a.one.Val, addTwoNumbers(p.one, p.two).Val)
	}
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

// linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	v, n := 0, 0

	for {
		v, n = add(l1, l2, n)
		temp.Val = v
		// 进入下一位链表节点
		l1 = next(l1)
		l2 = next(l2)
		// 如果两个数的下一位都为nil，则结束按位相加的运算
		if l1 == nil && l2 == nil { // 相加结束
			break
		}

		// 为下一位运算准备节点
		temp.Next = &ListNode{}
		temp = temp.Next
	}

	// n == 1 这是最后一次加运算进位，需要再添加一个节点
	if n == 1 {
		temp.Next = &ListNode{Val: n}
	}

	return result
}

// next  进入链表的下一位
func next(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

// v 是值， n是进位
func add(n1, n2 *ListNode, i int) (v, n int) {
	if n1 != nil {
		v += n1.Val
	}

	if n2 != nil {
		v += n2.Val
	}

	v += i

	if v > 9 {
		v -= 10
		n = 1
	}
	return
}
