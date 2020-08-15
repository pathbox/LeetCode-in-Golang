package LeetCode225

type MyStack struct {
	slice []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.slice = append(s.slice, x)
}

func (s *MyStack) Pop() int {
	if len(s.slice) == 0 {
		return -1
	}
	r := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return r
}

func (s *MyStack) Top() int {
	if len(s.slice) == 0 {
		return -1
	}
	return s.slice[len(s.slice)-1]
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.slice) == 0
}
