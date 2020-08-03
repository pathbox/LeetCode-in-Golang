package LeetCode055

func canJump(nums []int) bool {
	l := len(nums) - 1
	for i := l - 1; i >= 0; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0
}

/*
它更适合反过来想。我们可以从数组的倒数第二个个元素看起，每次往前遍历，如果当前的元素能够到达最后一个位置，那么我们把当前位置开始到最后全部“切断”，以当前元素为最后一个元素，重复第一次的过程就好了。而最后，如果我们遍历到数组的第一个元素，那么此题返回true，否则，你懂得
*/
