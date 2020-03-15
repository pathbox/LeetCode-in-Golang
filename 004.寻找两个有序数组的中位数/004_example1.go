package LeetCode004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	aLen := len(nums1)
	bLen := len(nums2)
	allLen := aLen + bLen
	i, j := 0, 0
	var tmpNums []int
	if allLen%2 == 1 { // 表示总数组是偶个数
		for {
			if i < aLen && j < bLen {
				if nums1[i] <= nums2[j] {
					tmpNums = append(tmpNums, nums1[i]) // 比较两个数组大小，将两个数组从小到大归并到一起
					i++
				} else {
					tmpNums = append(tmpNums, nums2[j])
					j++
				}
			} else if i >= aLen && j < bLen {
				tmpNums = append(tmpNums, nums2[j])
				j++
			} else if j >= bLen && i < aLen {
				tmpNums = append(tmpNums, nums1[i])
				i++
			}
			if len(tmpNums) == allLen/2+1 {
				return float64(tmpNums[allLen/2+1-1])
			}
		}
	} else {
		for {
			if i < aLen && j < bLen {
				if nums1[i] <= nums2[j] {
					tmpNums = append(tmpNums, nums1[i])
					i++
				} else {
					tmpNums = append(tmpNums, nums2[j])
					j++
				}
			} else if i >= aLen && j < bLen {
				tmpNums = append(tmpNums, nums2[j])
				j++
			} else if j >= bLen && i < aLen {
				tmpNums = append(tmpNums, nums1[i])
				i++
			}
			if len(tmpNums) == allLen/2+1 {
				return float64((tmpNums[allLen/2+1-1] + tmpNums[allLen/2-1])) / 2.0
			}
		}
	}
	return 0
}
