package LeetCode1714

func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	if k == 0 {
		return []int{}
	}

	// 求最小k个数，先建一个大小为k的大根堆
	heap := arr[:k]
	// 从n/2处开始堆化
	for i := k / 2; i >= 1; i-- {
		heapify(heap, k, i)
	}

	for i := len(arr) - 1; i >= k; i-- {
		// 若小于堆顶节点，则替换堆顶节点，并从根节点开始重新堆化
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			heapfy(heap, k, 1)
		}
	}
	return heap
}

// k表示堆的大小，i表示数组长度
func heapfy(arr []int, k, i int) {
	for i < k {
		max := i
		n := i * 2
		if n <= k && arr[i-1] < arr[n-1] {
			max = n
		}
		if n+1 <= k && arr[max-1] < arr[n] {
			max = n + 1
		}
		if max == i {
			break
		}
		swap := arr[i-1]
		arr[i-1] = arr[max-1]
		arr[max-1] = swap

		// +i
		i = max
	}
}
