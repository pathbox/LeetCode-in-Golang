package LeetCode658

func findClosestElements(arr []int, k int, x int) []int {
	start := 0
	end := len(arr) - 1

	removeNums := len(arr) - k
	for i := 0; i < removeNums; i++ {
		if x-arr[start] <= arr[end]-x {
			end--
		} else {
			start++
		}
	}
	return arr[start : end+1]
}

// O(n) O(1)
