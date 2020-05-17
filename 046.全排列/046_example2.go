package LeetCode046

func permute(nums []int) [][]int {

	res = [][]int{}
	for i := 0; i < len(nums); i++ {
		temp := make(map[int]bool)
		temp[i] = true
		l := make([]int, 0, len(nums))
		l = append(l, nums[i])
		dfs(nums, temp, l)
	}

	return res
}

var res [][]int

func dfs(nums []int, temp map[int]bool, l []int) {
	if len(l) == len(nums) {
		tt := make([]int, len(l))
		copy(tt, l)
		// fmt.Println("=====",res,l,tt)
		res = append(res, tt)
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := temp[i]; ok {
			continue
		}
		// fmt.Println(i,l)
		temp[i] = true
		l = append(l, nums[i])
		dfs(nums, temp, l)
		delete(temp, i)
		l = l[:len(l)-1]
	}

}
