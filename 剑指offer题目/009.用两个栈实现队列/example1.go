package Offer009

type CQueue struct {
	stack1 []int // 存储 是一个以stack1 的顶部为队列的head 底部为队列的tail
	stack2 []int // 临时迁移
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	for len(this.stack1) > 0 { // 不断的把stack1的数据从尾到头存储到stack2
		n := len(this.stack1)
		this.stack2 = append(this.stack2, this.stack1[n-1])
		this.stack1 = this.stack1[:n-1]
	}

	this.stack1 = append(this.stack1, value) // 再将新的值入栈stack1 此时stack1只有底部一个数value

	for len(this.stack2) > 0 { //再将stack2的原来stack1的数据导入回stack1
		n := len(this.stack2)
		this.stack1 = append(this.stack1, this.stack2[n-1])
		this.stack2 = this.stack2[:n-1]
	}
}

func (this *CQueue) DeleteHead() int {
	n := len(this.stack1)
	if n == 0 {
		return -1
	}
	defer func() {
		this.stack1 = this.stack1[:n-1]
	}()

	return this.stack1[n-1] // delete掉最后一个值,stack1栈顶的值是最早的，栈底的是最晚的
}
