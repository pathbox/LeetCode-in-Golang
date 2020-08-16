package LeetCode283

func moveZeroes(nums []int) {
	l := len(nums) //递减非零计数器
	i := 0         //递增非零计数器
	for i < l {
		if nums[i] == 0 {
			l--                                     //遇到一个0，减少1，最后与i比较作为结束条件
			nums = append(nums[0:i], nums[i+1:]...) //把0前后两部分合并 相当于移动了数组
			nums = append(nums, 0)                  //在末尾补回0
		} else {
			i++
		}
	}
}
