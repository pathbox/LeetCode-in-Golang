package LeetCode155

type MinStack struct {
	s []Node
}

type Node struct {
	d int // data
	m int // current min
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	d := Node{d: x, m: x}
	top := ms.s[len(ms.s)-1] //入栈前的顶部
	if len(ms.s) > 0 && top.m < x {
		d.m = top.m
	}
	ms.s = append(ms.s, d)
}

func (ms *MinStack) Pop() {
	// p := ms.s[len(ms.s)-1]
	ms.s = ms.s[:len(ms.s)-1]
}
func (ms *MinStack) Top() int {
	// p := ms.s[len(ms.s)-1]
	return ms.s[len(ms.s)-1].d
}

func (ms *MinStack) GetMin() int {
	return ms.s[len(ms.s)-1].m
}

// 有些空间换时间的意思，多了个m存当前的最小值
