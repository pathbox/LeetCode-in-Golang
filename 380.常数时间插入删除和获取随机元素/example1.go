package LeetCode380

import "math/rand"

type RandomizedSet struct {
	set    map[int]int // map
	data   []int       // slice
	length int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int, 100), make([]int, 0, 100), 0}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; !ok {
		this.data = append(this.data, val) // 从尾部加入
		this.set[val] = this.length
		this.length++
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.set[val]; ok {
		index := this.set[val]
		last_val := this.data[len(this.data)-1]
		this.data[index], this.data[len(this.data)-1] = this.data[len(this.data)-1], this.data[index] // 将要删除的元素和最后一个元素交换，然后再删除，这样就是常数时间删除了，不需要移动数组
		this.data = this.data[:len(this.data)-1]
		this.set[last_val] = index
		delete(this.set, val)
		this.length--
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.length > 0 {
		key := rand.Intn(this.length)
		return this.data[key]
	}
	return -1
}
