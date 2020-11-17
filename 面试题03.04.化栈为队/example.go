package LeetCode0304

type MyQueue struct {
	nums []int
	help []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		nums: make([]int, 0),
		help: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.nums = append(this.nums, x)
}

/** Removes the element from in front of queue and returns that element. */
// 在pop的时候，要先把nums中的数据都先传到help，再从help中栈出即可
func (this *MyQueue) Pop() int {
	this.Peek() // 先调用Peek保证help非空
	v := this.help[len(this.help)-1]
	this.help = this.help[:len(this.help)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.help) == 0 {
		for len(this.nums) != 0 {
			this.help = append(this.help, this.nums[len(this.nums)-1])
			this.nums = this.nums[:len(this.nums)-1]
		}
	}
	return this.help[len(this.help)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.nums) == 0 && len(this.help) == 0
}
