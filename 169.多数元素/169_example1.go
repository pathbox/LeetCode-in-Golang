package LeetCode169

func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return major
}

// 摩尔投票
// 为什么这里可以用摩尔投票呢？因为有一个条件是：多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 比如：现在输入的是 [3,2,3,1,6] 3是最多出现的次数是2，但是n/2 是2， 2 没有大于n/2所以得到的最后结果不是3，而是6
