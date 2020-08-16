package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func topKFrequent(nums []int, k int) []int {
	// 使用字典，统计每个元素出现的次数，元素为键，元素出现的次数为值
	hashMap := make(map[int]int, len(nums))
	for _, num := range nums {
		hashMap[num]++
	}
	// 遍历map，用最小堆保存频率最大的k个元素
	h := make(PriorityQueue, 0)
	//i := 0
	heap.Init(&h)

	for value, priority := range hashMap {
		if h.Len() < k {
			heap.Push(&h, &Item{
				value:    strconv.Itoa(value),
				priority: priority,
				//index:    i,
			})
		} else if priority > h[0].priority {
			heap.Pop(&h)
			heap.Push(&h, &Item{
				value:    strconv.Itoa(value),
				priority: priority,
				//index:    i,
			})
		}
	}
	// 取出最小堆中的元素
	var res []int
	for h.Len() > 0 {
		tmp := heap.Pop(&h).(*Item).value
		tmp1, _ := strconv.Atoi(tmp)
		res = append(res, tmp1)
	}
	return res
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	//index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority // 注意这里是需要修改的
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	//pq[i].index = i
	//pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	//n := len(*pq)
	item := x.(*Item)
	//item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//// update modifies the priority and value of an Item in the queue.
//func (pq *PriorityQueue) update(item *Item, value string, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}
