package Offer051

// 归并排序
func reversePairs(nums []int) int {
	var count int
	sortArray(nums, &count)
	return count
}

// 下面就是常规归并排序
// 分治
func sortArray(nums []int, cnt *int) []int {
	var lens = len(nums)
	if lens <= 1 {
		return nums
	}
	return Merge(sortArray(nums[:lens/2], cnt), sortArray(nums[lens/2:], cnt), cnt)
}

func Merge(Pre, Post []int, cnt *int) []int {
	var lenPre, lenPost = len(Pre), len(Post)
	var list []int
	var i, j int
	// 比较+合并
	for i < lenPre && j < lenPost {
		if Pre[i] <= Post[j] {
			list = append(list, Pre[i])
			i++
		} else {
			list = append(list, Post[j])
			// 这边处理下逆序数
			// 因为 post 切片里第 j 个数，比 pre 切片里 i 往后的数小 ，所以逆序数为 pre - i
			*cnt = *cnt + lenPre - i // Pre数组还没有归并的数 lenPre - i
			j++
		}
	}
	// 多出来数的直接接到后面
	if i < lenPre {

		list = append(list, Pre[i:]...)
	}
	if j < lenPost {
		list = append(list, Post[j:]...)
	}
	return list
}
