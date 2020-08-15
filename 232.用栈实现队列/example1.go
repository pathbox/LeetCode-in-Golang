package LeetCode232

// 不断的push到input栈中，input会栈的方式同步到output，所以，pop时候，取最尾部的数据则是最开始push的数据，这样就满足队列熟悉
type MyQueue struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) { // 不断的到input栈中
	this.input = append(this.input, x)
}

func (this *MyQueue) Pop() int {
	this.Peek()
	e := this.output[len(this.output)-1] // 取output的最后一个
	this.output = this.output[:len(this.output)-1]
	return e
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 { // output 为空的时候才要再将input的数据以栈的方式同步到output
		for len(this.input) > 0 {
			this.output = append(this.output, this.input[len(this.input)-1])
			this.input = this.input[:len(this.input)-1]
		}
	}
	return this.output[len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.input) == 0 && len(this.output) == 0 {
		return true
	}
	return false
}
