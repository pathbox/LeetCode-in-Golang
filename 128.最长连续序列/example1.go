package LeetCode128

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num, _ := range numSet {
		if !numSet[num-1] { // 为什么可以这样，只考虑递增的情况，因为如果满足条件，肯定会有某一个元素是该子序列的最小值，我们只需要找到这一种情况，然后递增的计算即可，如果这一个元素都没有，说明没有满足的情况
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
