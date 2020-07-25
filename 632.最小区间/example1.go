package LeetCode632

// https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/solution/zui-xiao-qu-jian-by-coldme-2/

// https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/solution/wo-ye-bu-zhi-dao-shi-sha-fang-fa-by-_days/
func smallestRange(nums [][]int) []int {
	r := [][]int{}
	totalCount := 0
	for _, v := range nums {
		totalCount += len(v)
	}
	c := make([]int, len(nums))
	for i := 0; i < totalCount; i++ {
		a := make([]int, 2)
		a[0] = 100001
		minIdx := 0
		for idx, v := range nums {
			if v[c[idx]] < a[0] {
				a[0] = v[c[idx]]
				minIdx = idx
			}
			if v[c[idx]] > a[1] {
				a[1] = v[c[idx]]
			}
		}
		r = append(r, a)
		if c[minIdx] < len(nums[minIdx])-1 {
			c[minIdx]++
		} else {
			break
		}
	}

	s := [][]int{}
	for _, v := range r {
		max, min := -100001, 100001
		for _, num := range v {
			if max < num {
				max = num
			}
			if min > num {
				min = num
			}
		}
		a := []int{min, max}
		s = append(s, a)
	}

	result := []int{}
	for _, v := range s {
		if len(result) == 0 {
			result = v
		} else {
			if v[1]-v[0] < result[1]-result[0] {
				result = v
			}
		}
	}
	return result
}
