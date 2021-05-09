package LeetCode325

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
