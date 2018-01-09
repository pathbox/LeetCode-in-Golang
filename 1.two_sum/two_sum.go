package Problem001

func twoSum(nums []int, target int) []int {
	// m: m[值]值的index
	m := make(map[int]int, len(nums))

	// i 是数组下标，b数元素值
	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		// a的序号是i，b的序号是j
		if j, ok := m[target-b]; ok { // target - b为a，不断的判断a在不在构造的map中，如果不存在，将a存入map，等待之后的比较
			// ok 为 true
			// 说明在i之前，存在nums[j] == a
			return []int{j, i}
			// 注意，顺序是j，i，因为j<i
		}

		// 如果没有找到，把i和i的值，存入map，构造map。map一开始是空的。
		m[nums[i]] = i
	}

	return nil

}
