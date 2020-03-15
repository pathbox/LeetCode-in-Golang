package LeetCode004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var newArray []int
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			newArray = append(newArray, nums1[i])
			i++
		} else if nums1[i] == nums2[j] {
			newArray = append(newArray, nums1[i])
			newArray = append(newArray, nums2[j])
			i++
			j++
		} else {
			newArray = append(newArray, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		newArray = append(newArray, nums1[i:len(nums1)]...)
	}
	if j < len(nums2) {
		newArray = append(newArray, nums2[j:len(nums2)]...)
	}
	length := len(newArray)
	if len(newArray)%2 == 0 { // 数组长度偶数
		return float64((newArray[length/2-1] + newArray[length/2])) / 2
	} else { // 数组长度奇数
		return float64(newArray[length/2])
	}
}

// 归并两个有序数组为一个数组后，根据数组元素个数的奇数偶数来返回中位数
