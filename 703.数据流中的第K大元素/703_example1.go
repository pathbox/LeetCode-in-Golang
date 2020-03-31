package LeetCode703

import "container/heap"

type KthLargest struct {
	k       int
	minHeap *IntMinHeap // 最小堆
}

func Constructor(k int, nums []int) KthLargest {
	ktl := KthLargest{minHeap: &IntMinHeap{}, k: k}
	for _, n := range nums {
		heap.Push(ktl.minHeap, n)
		if ktl.minHeap.Len() > k {
			heap.Pop(ktl.minHeap) //只保留k个元素，第k大就是[0]的元素
		}
	}
	return ktl
}

func (ktl *KthLargest) Add(val int) int {
	heap.Push(ktl.minHeap, val)
	if ktl.minHeap.Len() > ktl.k {
		heap.Pop(ktl.minHeap) //只保留k个元素，第k大就是[0]的元素
	}
	return ktl.minHeap.Peek() //第k大就是[0]的元素
}

// 小顶堆
type IntMinHeap []int

func (h *IntMinHeap) Len() int {
	return len(*h)
}

func (h *IntMinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n] // 排除了最后一个
	return x
}

func (h *IntMinHeap) Peek() int {
	return (*h)[0]
}
