package LeetCode0305

type stack struct {
	nums []int
}

func newStack() *stack {
	return &stack{
		nums: []int{},
	}
}

// 栈的操作
func (s *stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	r := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return r
}
func (s *stack) top() int {
	if s.isEmpty() {
		return -1
	}
	var res = s.nums[len(s.nums)-1] //取出第一个
	return res
}

func (s *stack) isEmpty() bool {
	if len(s.nums) == 0 {
		return true
	}
	return false
}

type SortedStack struct {
	nums *stack
	help *stack // 辅助栈
}

func Constructor() SortedStack {
	return SortedStack{
		nums: newStack(),
		help: newStack(),
	}
}

func (this *SortedStack) Push(val int) {
	for {
		if this.nums.isEmpty() || val <= this.nums.top() {
			break
		}
		this.help.push(this.nums.top())
		this.nums.pop()
	}
	this.nums.push(val)
	for {
		if this.help.isEmpty() {
			break
		}
		this.nums.push(this.help.top())
		this.help.pop()
	}
}
func (this *SortedStack) Pop() {

	if this.nums.isEmpty() {
		return
	}
	this.nums.pop()
}

func (this *SortedStack) Peek() int {
	if this.nums.isEmpty() {
		return -1
	}
	return this.nums.top()
}

func (this *SortedStack) IsEmpty() bool {
	return this.nums.isEmpty()
}
