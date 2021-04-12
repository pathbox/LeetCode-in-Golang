package LeetCode560

// 前缀和问题
func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1} // 0 空数组的情况
	var sum int
	var count int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}

/*
m存的是每个以i结尾的前缀和的次数,存在这种情况
m[0..j的前缀和]+m[j..i]=m[0..i的前缀和]
把m[j..i]看成是k，则 m[0..i的前缀和]-k= m[0..j的前缀和]
算法就是: 0.。i的前缀和减去k得到的值,是否也是某一个前缀和,到m中查找匹配。如果说是，说明 j..i这段连续的数字之和等于k，找到一个结果
*/
