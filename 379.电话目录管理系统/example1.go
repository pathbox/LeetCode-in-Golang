package LeetCode397

/*
队列存储电话号码，get 出队列， release 则加队列
map 存储是否已经使用
注意出队列的时候，判断一下当前是否存在在map中已经被使用
*/

type PhoneDirectory struct {
	queue  []int
	canuse map[int]bool // false 表示已经使用，true表示还未Get使用
}

func Constructor(maxNumbers int) PhoneDirectory {
	var phoneDirectory PhoneDirectory
	phoneDirectory.queue = make([]int, maxNumbers)
	phoneDirectory.canuse = make(map[int]bool, maxNumbers)
	for i := 0; i < maxNumbers; i++ {
		phoneDirectory.queue[i] = i
		phoneDirectory.canuse[i] = true
	}
	return phoneDirectory
}

func (this *PhoneDirectory) Get() int {
	if len(this.queue) > 0 {
		num := this.queue[0]
		this.canuse[num] = false
		this.queue = this.queue[1:] // 去掉queue[0]
		return num
	}
	return -1
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	return this.canuse[number]
}

func (this *PhoneDirectory) Release(number int) { // 将已用的归还置为可用
	if !this.canuse[number] {
		this.queue = append(this.queue, number)
		this.canuse[number] = true
	}
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */
