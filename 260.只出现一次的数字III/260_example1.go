package LeetCode260

/*
1.通过所有数据的异或resultNum--->单个出现异或结果集
2.设置标志位diff----->用于区分两个单体数据
3.求出一个单体数据x---->再用x^resultNum求出另一个
*/

func singleNumber(nums []int) []int {
	var resultNum int
	for i := 0; i < len(nums); i++ {
		resultNum = resultNum ^ nums[i] // 得到的是两个数字的异或值
	}
	var diff = resultNum & (-resultNum) // 得到某一位是1，以此来进行分组
	var x int
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] != 0 {
			x = x ^ nums[i] // 得到某一组中的单个数字
		}
	}
	return []int{x, x ^ resultNum} // 将某x与之前的结果再异或，能得到另一个数
}

// 时间复杂度：O(n) ,空间复杂度：O(1)
