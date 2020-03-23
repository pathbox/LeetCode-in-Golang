package LeetCode088

// 按最大值从后往前赋值,先比较懂得值，从大的值开始合并
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 { // n是nums2的长度
		if m <= 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
