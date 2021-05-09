package LeetCode325

/*
n 的前缀和 S(n)S(n) 表示 0 \sim n0∼n 所有元素的和。

对于 nn，我们如果要找出 nn 及 nn 前面元素组成的和为 kk 的最长子数组，而 S(n)S(n) 和 kk 之间的差为 S(n) - kS(n)−k。此时如果 nn 前面存在元素 mm 的前缀和 S(m) == S(n) - kS(m)==S(n)−k，也就是 S(n)-S(m)==k, 那么 m + 1 ∼n 元素即为我们所求的最长子数组(这段和为k),。如果不存在元素 m，则 n 前面暂时无法组成和为 k 的最长子数组

*/

func maxSubArrayLen(nums []int, k int) int {
	//前缀和字典，key是前缀和，v是最小的索引
	m := make(map[int]int)
	var res int
	var sum int
	//初始值是-1，这时因为比较的是长度，第一个值索引为0，包含第一个的长度应该再+1，初始值索引为-1方便比较
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		//再判断sum-k存不存在
		if v, ok := m[sum-k]; ok {
			res = max(res, i-v)
		}
		//再存储sum，本题遇到相同值只存储第一个sum的下标，因为要求的是最大，所以存后面的没有意义
		// 如果说是想要找最短长度的话， 这里需要为ok，表示索引不断被更大的更新
		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
