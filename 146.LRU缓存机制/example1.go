package LeetCode146

/*
首先是通过双向链表来存放数据，插入和删除实现O(1)。分别记录链表的header tail指针
对于查找链表是无法实现O(1)，需要借助于map来实现
为了在链表操作过程中更少的判断头尾节点是否为null，采用哨兵机制，头尾都添加哨兵.
这里只是实现一个线程不安全的版本
*/

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

// 非安全版，安全需要加锁,先操作链表，再操纵map
func Constructor(cap int) LRUCache {
	cache := LRUCache{
		cap:    cap,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, cap),
	}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		this.remove(node) // 链表中删了这个节点，再新建一个节点放入头部
		this.putHead(node)
		return this.m[key]
	}
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.m[key]; ok {
		//　存在节点.更新数据。移动到head
		node.value = value
		this.remove(node)
		this.putHead(node)
	} else {
		// 不存在
		// 如果容器已经满了，需要删除链表尾部,从map中删除
		if len(this.m) >= this.cap {
			delKey := this.tail.pre.key
			this.remove(this.tail.pre)
			delete(this.m, delKey)
		}
		//创建新的节点,并放在head,添加到map中
		newNode := Node{key: key, value: value}
		this.putHead(&newNode)
		this.m[key] = &newNode
	}
}

func (this *LRUCache) putHead(node *Node) {
	originNext := this.header.next
	this.header.next = node
	node.next = originNext
	originNext.pre = node
	node.pre = this.header
}

// 因为是链表节点，本身是指针，能够定位位置，不像数组需要遍历需要，所以这里的复杂度是O(1)
func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
