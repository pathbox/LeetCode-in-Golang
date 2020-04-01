package LeetCode705

import "strconv"

/*
基于 HashMap 设计理念的 Int Set
*/

const (
	defaultSize      int     = 10
	growLoadFactor   float32 = 0.75
	shrinkLoadFactor float32 = 0.25
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func newNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

func newDefaultSlice(size, defaultVal int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = defaultVal
	}
	return arr
}

func Constructor() MyHashSet {
	return MyHashSet{
		bucketArr: make([]*Node, defaultSize),
	}
}

func djb2Hash(key int) uint {
	bs := []byte(strconv.Itoa(key))
	h := uint(5381)
	for _, r := range bs {
		h = (h << 5) + h + uint(r)
	}
	return h
}

func getSlot(key, size int) int {
	return int(djb2Hash(key)) % size
}

type MyHashSet struct {
	bucketArr []*Node
	size      int
}

func (t *MyHashSet) loadFactor() float32 {
	return float32(t.size) / float32(len(t.bucketArr))
}

func (t *MyHashSet) resize(size int) {
	newBucket := make([]*Node, size)
	for i := range t.bucketArr {
		node := t.bucketArr[i]
		var keysInBucket []int
		for {
			if node == nil {
				break
			}
			keysInBucket = append(keysInBucket, node.Val)
			node = node.Next
		}

		for j := range keysInBucket {
			key := keysInBucket[j]
			newSlot := getSlot(key, size)
			node = newBucket[newSlot]
			head := newNode(key)
			head.Next = node
			if node != nil {
				node.Prev = head
			}
			newBucket[newSlot] = head
		}
	}
	t.bucketArr = newBucket
}

func (t *MyHashSet) grow() {
	newSize := len(t.bucketArr) * 2
	t.resize(newSize)
}

func (t *MyHashSet) shrink() {
	newSize := len(t.bucketArr) / 2
	if defaultSize >= newSize {
		newSize = defaultSize
	}
	t.resize(newSize)
}

func (t *MyHashSet) searchNode(key int) (*Node, bool) {
	slot := getSlot(key, len(t.bucketArr))
	node := t.bucketArr[slot]

	for {
		switch {
		case node == nil:
			fallthrough
		case node.Next == nil && node.Val != key:
			return node, false
		case node.Val == key:
			return node, true
		default:
			node = node.Next
		}
	}
}

func (t *MyHashSet) Add(key int) {
	node, ok := t.searchNode(key)
	if ok {
		return
	}
	newKey := newNode(key)
	if node != nil {
		node.Next = newKey
		newKey.Prev = node
	} else {
		slot := getSlot(key, len(t.bucketArr))
		t.bucketArr[slot] = newKey
	}

	t.size++

	if t.loadFactor() >= growLoadFactor {
		t.grow()
	}
}

func (t *MyHashSet) Remove(key int) {
	node, ok := t.searchNode(key)
	if !ok {
		return
	}
	slot := getSlot(key, len(t.bucketArr))
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		t.bucketArr[slot] = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	t.size--

	if t.loadFactor() <= shrinkLoadFactor {
		t.shrink()
	}
}

/** Returns true if this set contains the specified element */
func (t *MyHashSet) Contains(key int) bool {
	_, ok := t.searchNode(key)
	return ok
}
