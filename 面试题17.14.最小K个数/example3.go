package LeetCode1714

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	return sort(arr, 0, len(arr)-1, k)
}

func sort(arr []int, start, end, k int) []int {
	idx := partition(arr, start, end)
	if idx+1 == k {
		return arr[:idx+1]
	} else if idx+1 < k { // idx+1更小，从 idx+1:end中继续找k
		return sort(arr, idx+1, end, k)
	} else {
		return sort(arr, start, idx-1, k)
	}
}

func partition(arr []int, start, end int) int {
	privot := end
	l, r := start, end-1
	for l <= r {
		for l <= r && arr[l] <= arr[pivot] {
			l++
		}
		for l <= r && arr[r] >= arr[pivot] {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	arr[l], arr[pivot] = arr[pivot], arr[l]
	return l // you got the partition index
}
