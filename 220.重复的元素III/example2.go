package LeetCode220

/*
	桶排序
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	buckets := make(map[int]int)
	w := t + 1
	for i := range nums {
		key := keyFunc(nums[i], w)
		if _, ok := buckets[key]; ok {
			return true
		}
		if val, ok := buckets[key-1]; ok && abs(nums[i]-val) < w {
			return true
		}
		if val, ok := buckets[key+1]; ok && abs(nums[i]-val) < w {
			return true
		}
		buckets[key] = nums[i]
		if i >= k {
			delete(buckets, keyFunc(nums[i-k], w))
		}
	}
	return false
}

func abs(data int) int {
	if data < 0 {
		return -1 * data
	}
	return data
}

func keyFunc(x, w int) int {
	if x < 0 {
		return (x+1)/w - 1
	} else {
		return x / w
	}
}
