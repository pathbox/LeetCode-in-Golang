/* 设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。
get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
put(key, value) - 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。
进阶：
你是否可以在 O(1) 时间复杂度内执行两项操作？

不是只有一个双向链表，而是一个双向链表map
*/

package LeetCode705

type LFUCache struct {
	cache               map[int]*Node       // 缓存层，读的时候O(1)
	freq                map[int]*DoubleList // 优先级层，优先级和过滤层
	ncap, size, minFreq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache: make(map[int]*Node),
		freq:  make(map[int]*DoubleList),
		ncap:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.IncFreq(node) // 经常被使用的key，freq会不断的增加
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key, value int) {
	if this.ncap == 0 {
		return
	}
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.IncFreq(node)
	} else {
		if this.size >= this.ncap {
			node := this.freq[this.minFreq].RemoveLast() // 从minFreq层,移除最后的节点,过滤调不经常使用的节点
			delete(this.cache, node.key)
			this.size--
		}
		x := &Node{key: key, val: value, freq: 1} // 每次数据都优先存到 freq = 1这一层的双向链表中
		this.cache[key] = x
		if this.freq[1] == nil {
			this.freq[1] = CreateDL()
		}
		this.freq[1].AddFirst(x)
		this.minFreq = 1
		this.size++
	}
}

// 将热freq的递增
func (this *LFUCache) IncFreq(node *Node) {
	_freq := node.freq                                       // node.freq 节点所在的链表map
	this.freq[_freq].Remove(node)                            // 先删除node的老位置
	if this.minFreq == _freq && this.freq[_freq].IsEmpty() { // 如果这个freq map值是空的，从map中删除
		this.minFreq++
		delete(this.freq, _freq)
	}

	node.freq++                      // 递增
	if this.freq[node.freq] == nil { // 递增后位置对应的链表还是空的，初始化一个链表
		this.freq[node.freq] = CreateDL()
	}
	this.freq[node.freq].AddFirst(node) // 将节点加入到该链表的头部
}

type DoubleList struct {
	head, tail *Node
}

type Node struct {
	prev, next     *Node
	key, val, freq int
}

// 初始化一个双向链表
func CreateDL() *DoubleList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return &DoubleList{
		head: head,
		tail: tail,
	}
}

// 操作双向链表，加入到链表头部
func (this *DoubleList) AddFirst(node *Node) {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

// 链表中删除节点
func (this *DoubleList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil // 之后这个node就会被释放
	node.next = nil
}

// 删除最后一个节点
func (this *DoubleList) RemoveLast() *Node {
	if this.IsEmpty() {
		return nil
	}

	last := this.tail.prev
	this.Remove(last)

	return last
}

// 双向链表是否为空
func (this *DoubleList) IsEmpty() bool {
	return this.head.next == this.tail
}
