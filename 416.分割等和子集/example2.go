package LeetCode416

func canPartition(nums []int) bool {
	total := 0
	m := make(map[int]int)
	for _, n := range nums {
		total += n
		m[n]++
	}
	//合为奇数则直接为false
	if total%2 == 1 {
		return false
	}

	return canPartitionWithMap(total/2, m)
}

func canPartitionWithMap(target int, m map[int]int) bool {
	if target == 0 {
		return true
	}
	if len(m) == 0 {
		return false
	}
	for k, _ := range m {
		m[k]--
		if m[k] == 0 {
			delete(m, k)
		}
		if target-k >= 0 && canPartitionWithMap(target-k, m) {
			return true
		}
		m[k]++
	}
	return false
}
