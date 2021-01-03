package LeetCode480

import (
	"container/list"
	"sort"
)

// 解法一 用链表按照题意实现 时间复杂度 O(n * k) 空间复杂度 O(k) 维护有序的list窗口
func medianSlidingWindow(nums []int, k int) []float64 {
	var res []float64
	w := getWindowList(nums[:k], k)
	res = append(res, getMedian(w, k))

	for p1 := k; p1 < len(nums); p1++ {
		w = removeFromWindow(w, nums[p1-k]) // 相当于left++
		w = insertInWindow(w, nums[p1])     // 相当于right++
		res = append(res, getMedian(w, k))
	}
	return res
}

func getWindowList(nums []int, k int) *list.List {
	s := make([]int, k)
	copy(s, nums)
	sort.Ints(s)
	l := list.New()
	for _, n := range s {
		l.PushBack(n)
	}
	return l
}

func removeFromWindow(w *list.List, n int) *list.List {
	for e := w.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == n {
			w.Remove(e)
			return w
		}
	}
	return w
}

func insertInWindow(w *list.List, n int) *list.List {
	for e := w.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= n {
			w.InsertBefore(n, e)
			return w
		}
	}
	w.PushBack(n)
	return w
}

func getMedian(w *list.List, k int) float64 {
	e := w.Front()
	for i := 0; i < k/2; e, i = e.Next(), i+1 {
	}
	if k%2 == 1 {
		return float64(e.Value.(int))
	}
	p := e.Prev()
	return (float64(e.Value.(int)) + float64(p.Value.(int))) / 2
}
