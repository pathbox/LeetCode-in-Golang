package LeetCode211

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: newTrieNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]] // trie树的遍历
	}
	node.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this.root
	return search(word, node)
}

func search(word string, node *TrieNode) bool {
	if node == nil {
		return false
	}
	if word == "" {
		return node.isEnd // 能够遍历到最后一个node说明search成功
	}

	if word[0] == '.' {
		for _, n := range node.children {
			if search(word[1:], n) {
				return true
			}
		}
	}
	node = node.children[word[0]]
	return search(word[1:], node)
}
