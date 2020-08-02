package LeetCode018

func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	// 生序排序
	sort.Ints(nums)
	// 遍历、收集等于targe的组合
	var res [][]int
	for i := 0; i < len(nums)-3; i++{
		// 第一个元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < len(nums)-2; j++{
			// 第二个元素去重
			if j >i+1 && nums[j] == nums[j-1]; {
				continue
			}

			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i]+nums[j]+nums[k]+nums[l]
				x := sum -target
				if x == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
				}
				if x > 0 {
					// 第四个元素去重,并往小的方向移动
					for cur := l; l > k && nums[l] == nums[cur];l--{
						continue
					}
				}else {
					// 第三个元素去重,并往小的方向移动
					for cur := k; k < l && nums[k] == nums[cur]; k--{
						continue
					}
				}
			}
		}
	}
	return res
}