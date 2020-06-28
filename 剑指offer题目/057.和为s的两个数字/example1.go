package Offer057

func twoSum(nums []int, target int) []int {
	m := make(map[int]bool, 0)
	for i, _ := range nums {
		m[nums[i]] = true
	}
	l := len(nums)
	for i := 0; i < l; i++ {
		a := nums[i]
		x := target - a
		if m[x] {
			return []int{a, x}
		}
	}
	return nil
}
