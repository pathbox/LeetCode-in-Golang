package LeetCode078

var (
	res [][]int
)

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	backtrack([]int{}, nums, 0)
	return res
}

func backtrack(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtrack(temp, nums, i+1)

		// 通过删掉最后一个元素实现回溯， 才会有 回溯 1， 再回溯 2，最后回溯 3
		temp = temp[:len(temp)-1]
	}
}
