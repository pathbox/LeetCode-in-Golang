package LeetCode384

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

// /** Returns a random shuffling of the array. */
// func (this *Solution) Shuffle() []int {
// 	nums := make([]int, len(this.nums))
// 	buf := make([]int, len(this.nums))
// 	copy(buf, this.nums)
// 	for i := range nums {
// 		j := rand.Intn(len(buf))
// 		nums[i] = buf[j]
// 		buf = append(buf[0:j], buf[j+1:]...)
// 	}
// 	return nums
// }
