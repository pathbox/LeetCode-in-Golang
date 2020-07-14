package LeetCode219

/*
1、维护一个哈希表m，存储每个数字上次出现的位置
2、遍历nums
3、如果nums[i]这个数字，上次出现过，且，出现的位置和现在距离小于k，可以直接返回true
4、遍历结束后，还没出现3中的情况，那么就返回false吧
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for idx, v := range nums {
		index, ok := m[v] // idx >= index
		if ok && idx-index <= k {
			return true
		}
		m[v] = idx
	}
	return false
}
