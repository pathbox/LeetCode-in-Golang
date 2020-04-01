package LeetCode705

type MyHashSet struct {
	set []byte
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		set: make([]byte, 1000),
	}
}

func (this *MyHashSet) Add(key int) {
	if key/8 >= len(this.set) {
		expandLen := ((key / 8) - len(this.set)) + 1
		if expandLen > 1000000/8+1 {
			expandLen = 1000000/8 + 1
		}
		this.set = append(this.set, make([]byte, expandLen)...) // 对set进行扩容
	}
	this.set[key/8] = (1 << (key % 8)) | this.set[key/8]
}

func (this *MyHashSet) Remove(key int) {
	if key/8 < len(this.set) {
		this.set[key/8] = ^(1 << (key % 8)) & this.set[key/8]
	}
}

func (this *MyHashSet) Contains(key int) bool {
	exist := false
	if key/8 >= len(this.set) {
		exist = false
	} else if (1<<(key%8))&this.set[key/8] == 1<<(key%8) {
		exist = true
	}
	return exist
}
