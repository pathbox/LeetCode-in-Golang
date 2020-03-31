package LeetCode705

type MyHashSet struct {
	Set []*Bucket
}

const L = 769

func Constructor() MyHashSet {
	hs := MyHashSet{}
	for i := 0; i < L; i++ {
		hs.Set = append(hs.Set, NewBucket())
	}
	return hs
}

func Hash(key int) int {
	return key % L
}

func (hs *MyHashSet) Add(key int) {
	bucketIndex := Hash(key)
	hs.Set[bucketIndex].insert(key)
}

func (hs *MyHashSet) Contains(key int) bool {
	bucketIndex := Hash(key)
	return hs.Set[bucketIndex].exists(key)
}

func (hs *MyHashSet) Remove(key int) {
	bucketIndex := Hash(key)
	hs.Set[bucketIndex].delete(key)
}

type Node struct {
	Val  int
	Next *Node
}

type Bucket struct {
	Head *Node
}

func NewBucket() *Bucket {
	return &Bucket{
		&Node{0, nil},
	}
}

func newNode(key int) *Node {
	return &Node{
		key,
		nil,
	}
}

func (b *Bucket) insert(key int) {
	cur := b.Head.Next
	Niobe := newNode(key)
	for cur != nil {
		if cur.Val != key {
			if cur.Next != nil {
				cur = cur.Next
			} else {
				cur.Next = Niobe
				break
			}
		} else {
			break
		}
	}
	if cur == nil {
		b.Head.Next = Niobe
	}
}

func (b *Bucket) delete(key int) {
	cur := b.Head
	for cur.Next != nil {
		if cur.Next.Val == key {
			cur.Next = cur.Next.Next
			break
		} else {
			cur = cur.Next
		}
	}
}

func (b *Bucket) exists(key int) bool {
	cur := b.Head.Next
	for cur != nil {
		if cur.Val == key {
			return true
		} else {
			cur = cur.Next
		}
	}
	return false
}
