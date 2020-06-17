package Offer040

import "container/heap"

// 当大顶堆元素长度为k时，遇到小于根节点的值则弹出堆根元素插入该元素
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(x, y int) bool {
	return h[x] > h[y]
}

func (h IntHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *IntHeap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	intHeap := make(IntHeap, 0, k)
	heap.Init(&intHeap)
	for _, v := range arr {
		if len(intHeap) < k {
			heap.Push(&intHeap, v)
		} else if len(intHeap) == k {
			if v < intHeap[0] {
				heap.Pop(&intHeap)
				heap.Push(&intHeap, v)
			}
		}
	}
	return intHeap
}
