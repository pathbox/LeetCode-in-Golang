package LeetCode719

import "sort"

// 首先需要对数组排序，这样能很快找出最大距离对，也就是nums[len-1]-nums[0]，然后我们知道第k小的距离对一定会在最大和最小距离对之间，一般这种范围查找用到的就是二分查找，核心是中间那一段for循环代码，mid指的是中间的距离对的长度，left和right分别是nums上的左右指针，通过while循环来将所有<=mid的距离对的个数算出来，并加到count中，如果count>=k，那么right=mid，如果count<k，那么left=mid+1；这里是二分查找的常规模板（因为计算mid时除以2是向下取整的，所以count<k时，目标值一定在mid右边，反之count>=k，可能目标值就是mid或者在左边；如果对这点始终有疑问，可以去看看探索里的二分查找模板~）

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		mid := low + (high-low)>>1 // mid 是dis的二分中间
		if check(nums, mid, k) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func check(nums []int, dis int, k int) bool {
	low, count := 0, 0 // count 是 = dis 的区间大小
	for high, _ := range nums {
		for nums[high]-nums[low] > dis {
			low++
		}
		count += high - low // <=dis 的个数累加，每回有high - low个。因为low++，所以剩下的就是high - low
	}
	return count >= k
}

// O(n*(high-low)*logn)
