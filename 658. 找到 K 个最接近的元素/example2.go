package LeetCode658

/*
只要找到左边界，加上k，就可以得到右边界。
使用二分法，对于当前的mid，对比mid+k到x的距离，如果
1）x- arr[mid] > arr[mid+k]-x，说明左边界可以右移，left = mid+1
2）x- arr[mid] <= arr[mid+k]-x，则反之，right = mid
截取arr从left到left + k部分
*/
func findClosestElements(arr []int, k int, x int) []int {
	l := len(arr)
	left, right := 0, l-k
	for left < right {
		mid := left + (right-left)>>1
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
