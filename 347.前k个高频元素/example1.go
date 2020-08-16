package LeetCode347

import "container/heap"

/*
借助哈希表来建立数字和其出现次数的映射，遍历一遍数组统计元素的频率
维护一个元素数目为k的最小堆
每次都将新的元素与堆顶元素(堆中频率 最小的元素)进行比较
如果新的元素的频率比堆顶端的元素大，则弹出堆顶端的元素，将新的元素添加进堆中
最终，堆中的k个元素即为前k个高频元素
*/
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

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // 小顶堆
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
