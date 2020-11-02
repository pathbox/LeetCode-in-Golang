package main

import "errors"

// 二叉堆实现优先队列 --------------------------------------------------
type MyPriorityQueueInt struct {
	items []int
}

// Push 插入的时候，是插入到items末尾，会将插入的元素进行swim上浮操作
func (pq *MyPriorityQueueInt) swim(x int) {
	// 从下向上浮，只需要跟父节点比较，一般用在插入
	for f := x / 2; x > 1 && pq.items[x] > pq.items[f]; {
		// 父子交换
		pq.items[x], pq.items[f] = pq.items[f], pq.items[x]
		x = f
		f = x / 2
	}
}

func (pq *MyPriorityQueueInt) sink(x int) {
	// 从上向下沉
	for l, m := len(self.items)-1, 2*x; m <= l; m = 2 * x { // m会是一个小树中最大的索引
		// 在它和它的两个子节点中，寻找最大的那个跟它本身交换
		// 先跟左边的比,m指向更大的
		if pq.items[x] >= pq.items[m] {
			m = x
		}
		// 再跟右边的比,m指向更大的
		if n := x*2 + 1; n <= l {
			if pq.items[n] > pq.items[m] {
				m = n
			}
		}
		if m == x { // 最大值是本身的话，就停止下沉
			break
		} else { // 最大值是子节点，那就继续下沉
			pq.items[x], pq.items[m] = pq.items[m], pq.items[x]
			x = m // 继续下沉比较
		}
	}
}

func (pq *MyPriorityQueueInt) Push(x int) {
	if len(pq.items) == 0 {
		pq.items = append(pq.items, 0)
	}
	pq.items = append(pq.items, x)
	pq.swim(len(pq.items) - 1) // 对刚插入最后一个元素进行上浮操作
}

func (pq *MyPriorityQueueInt) Pop() (int, error) {
	if len(pq.items) == 0 {
		return 0, errors.New("Empty queue!")
	}
	x := pq.items[1]
	pq.items[1] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.sink(1) // 将pq.items[1]进行下沉操作
	return x, nil
}

func (pq *MyPriorityQueueInt) Length() int {
	return len(pq.items) - 1
}

func main() {

}
