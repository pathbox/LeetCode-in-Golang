package Offer030

type MinStack struct {
	nums []int //储存栈
	min  []int //辅助储存栈，存储最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0, 3),
		make([]int, 0, 3),
	}
}

// 每次push或pop都同时操作nums和min两个数组
func (this *MinStack) Push(x int) { // 每次Push的时候，也把一个对应的最小值push到min数组,min数组可能会有重复的最小值
	this.nums = append(this.nums, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else if this.min[len(this.min)-1] < x {
		this.min = append(this.min, this.min[len(this.min)-1])
	} else {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}
