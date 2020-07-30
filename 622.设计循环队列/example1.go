package LeetCode622

// 队空的判断条件是 head == tail,
// 队满的判断条件，(tail+1)%n=head。 尾指针+1 mod len 等于head 说明队列满了
// 当队列满时，tail 指向的位置实际上是没有存储数据的。所以，循环队列会浪费一个数组的存储空间。

type MyCircularQueue struct {
	Head  int
	Tail  int
	Value []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Head:  0,
		Tail:  0,
		Value: make([]int, k+1),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Value[this.Tail] = value
	this.Tail = (this.Tail + 1) % len(this.Value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = (this.Head + 1) % len(this.Value)
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Value[this.Head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	// (this.Tail-1+len(this.Value))%len(this.Value) 是队尾的指向的元素
	// 为了防止 -1 越界做的小处理
	return this.Value[(this.Tail-1+len(this.Value))%len(this.Value)]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.Head == this.Tail {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.Tail+1)%len(this.Value) == this.Head {
		return true
	}
	return false
}
