package LeetCode208

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Root = NewTrieNode()
	return t
}

func NewTrieNode() *TrieNode {
	n := &TrieNode{}
	n.Children = make(map[rune]*TrieNode)
	n.End = false
	return n
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.Root = NewTrieNode()
	return t
}

func (t *Trie) Insert(word string) {
	if len(word) < 1 {
		return
	}
	chars := []rune(word)
	slen := len(chars)
	node := t.Root
	for i := 0; i < slen; i++ {
		if _, exist := node.Children[chars[i]]; !exist {
			node.Children[chars[i]] = NewTrieNode()
		}
		node = node.Children[chars[i]]
	}
	node.End = true
}

func (t *Trie) Search(word string) bool {
	chars := []rune(word)
	slen := len(chars)
	node := t.Root
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; !exists {
			return false
		}
		node = node.Children[chars[i]]
		if node.End == true && i == slen-1 { // 二者刚好都到结尾,不然不断的从Children往下走
			return true
		}
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	chars := []rune(prefix)
	slen := len(chars)
	node := t.Root
	for i := 0; i < slen; i++ {
		if _, exists := node.Children[chars[i]]; !exists {
			return false
		}
		node = node.Children[chars[i]]
	}
	return true
}
