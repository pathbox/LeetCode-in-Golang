package LeetCode448

func findDisappearedNumbers(nums []int) []int {
	hash := map[int]int{}
	res := []int{}
	for _, i := range nums {
		hash[i] = 1
	}
	for j := 1; j <= len(nums); j++ {
		if hash[j] != 1 {
			res = append(res, j)
		}
	}
	return res

}
