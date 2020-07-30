package LeetCode973

import "container/heap"

// 最大堆
func kClosest(points [][]int, K int) [][]int {
	maxHeap := &PointMaxHeap{}
	for _, point := range points {
		heap.Push(maxHeap, point)
		if maxHeap.Len() > K {
			heap.Pop(maxHeap)
		}
	}
	res := make([][]int, 0, K)
	for maxHeap.Len() > 0 {
		res = append(res, heap.Pop(maxHeap).([]int))
	}
	return res
}

type PointMaxHeap [][]int

func (pq *PointMaxHeap) Len() int {
	return len(*pq)
}
func (pq *PointMaxHeap) Less(i, j int) bool {
	return (*pq)[i][0]*(*pq)[i][0]+(*pq)[i][1]*(*pq)[i][1] > (*pq)[j][0]*(*pq)[j][0]+(*pq)[j][1]*(*pq)[j][1]
}
func (pq *PointMaxHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PointMaxHeap) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PointMaxHeap) Pop() interface{} {
	n := len(*pq) - 1
	x := (*pq)[n]
	*pq = (*pq)[:n]
	return x
}

func (pq *PointMaxHeap) Peek() []int {
	return (*pq)[0]
}
