package Offer039

// 排序 因为重复值超过了一半 肯定中间值也是
func majorityElement(nums []int) int {
	m := make(map[int]int)
	t := len(nums) / 2
	for _, v := range nums {
		if m[v] >= t { //因为重复值超过了一半
			return v
		}
		m[v]++
	}
	return -1
}
