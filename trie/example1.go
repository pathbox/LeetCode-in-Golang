package main

import "fmt"

// type Trie struct {
// 	children map[byte]*Trie
// 	eow      bool
// }

// func Constructor() Trie {
// 	return Trie{children: map[byte]*Trie{}, eow: false} // 父节点是空值的
// }

// func (t *Trie) Insert(word string) {
// 	for i := 0; i < len(word); i++ {
// 		node, ok := t.children[word[i]]
// 		if !ok { // 如果这个节点不在trie中，则新建一个Trie节点，然后加到当前节点t的子节点中
// 			node = &Trie{
// 				children: map[byte]*Trie{}, eow: false,
// 			}
// 			t.children[word[i]] = node // 加到当前节点的子节点中
// 		}
// 		t = node // 将当前节点往下移
// 	}
// 	t.eow = true // 表示当前节点t是一个字符串的最后节点，可以用于匹配这个字符串
// }

// func (t *Trie) Search(word string) bool {
// 	for i := 0; i < len(word); i++ {
// 		node, ok := t.children[word[i]]
// 		if !ok {
// 			return false
// 		} // 在遍历整个word之前，中间有不满足条件的字符出现，则说明不完全匹配，返回
// 		t = node // 不断的往下走
// 	}
// 	return t.eow // 返回匹配到的最后一个节点的标记
// }

// // 匹配一个前缀，和 上面的Search方法有所不同。search表示这个word有在Tire中(从第一个到这个分支的最后一个节点刚好是这个word)，而匹配前缀表示 这一段前缀有被匹配
// func (t *Trie) StartsWith(prefix string) bool {
// 	for i := 0; i < len(prefix); i++ {
// 		node, ok := t.children[prefix[i]]
// 		if !ok {
// 			return false
// 		}
// 		t = node
// 	}
// 	return true
// }
type Trie struct {
	IsTerminated bool
	Children     map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	m := make(map[rune]*Trie)
	return Trie{IsTerminated: false, Children: m}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{IsTerminated: false, Children: make(map[rune]*Trie)}
			parent.Children[ch] = newChild
			parent = newChild
		}
	}
	parent.IsTerminated = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.IsTerminated
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.Children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

func main() {
	trieTree := Constructor()
	trieTree.Insert("apple")
	trieTree.Insert("banana")
	trieTree.Insert("/app/api/index")
	b1 := trieTree.Search("apple")
	b2 := trieTree.Search("banana")
	b3 := trieTree.Search("bana")
	b4 := trieTree.StartsWith("/app/")
	b5 := trieTree.StartsWith("/app/api/index")

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)
	fmt.Println("b4: ", b4)
	fmt.Println("b5: ", b5)
}
