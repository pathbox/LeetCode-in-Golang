package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// 使用字典，统计每个元素出现的次数，元素为键，元素出现的次数为值
	hashMap := make(map[int]int, len(nums))
	for _, num := range nums {
		hashMap[num]++
	}
	// 遍历map，用最小堆保存频率最大的k个元素
	var h IntHeap
	heap.Init(&h)

	for key, v := range hashMap {
		if h.Len() < k {
			heap.Push(&h, key)
		} else if v > hashMap[h[0]] {
			//heap.Pop(&h)
			//heap.Push(&h, key)
			heap.Fix(&h, key)
		}
	}
	// 取出最小堆中的元素
	var res []int
	for h.Len() > 0 {
		tmp := heap.Pop(&h).(int)
		res = append(res, tmp)
	}
	return res
}
func main() {
	nums := []int{3, 0, 1, 0}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
