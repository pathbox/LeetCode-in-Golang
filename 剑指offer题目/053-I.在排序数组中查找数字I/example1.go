package Offer053

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	// 分别查找left和right
	i := 0
	j := l - 1
	for i <= j {
		mid := j - (j-i)/2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	right := i
	i = 0
	j = l - 1
	for i <= j {
		mid := j - (j-i)/2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	left := j
	return right - left - 1
}

/*
另一种解法：
1. 二分法找到target
2. 当前对应index 一个往右一个往左继续比较，和target相等则累加1， 当不和target相等则停止。边界是0-len
*/
