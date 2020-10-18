package LeetCode239

// https://leetcode-cn.com/problems/sliding-window-maximum/solution/go-shuang-duan-jie-fa-zhu-shi-by-tian-fu-luo/

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}

	window, res := []int{}, []int{}
	for i, x := range nums {
		if i >= k && window[0] <= i-k { // window[0]已经进入了res，在这一轮去除
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= x { // 当前值x和window对应的索引的值进行比较，如果x更大，则window的索引出队列直到大于x的索引为止
			window = window[:len(window)-1]
		}
		window = append(window, i) // 将当前索引入队列,比x值小的索引都出了队列
		if i >= k-1 {              // 窗口是k-1
			res = append(res, nums[window[0]]) // window[0]存的是滑动窗口的最大值,保证window[0]存的是当前窗口最大值的索引
		}
	}
	return res
}
