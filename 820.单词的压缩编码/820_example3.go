package LeetCode820

// 使用字典树解决后缀重复（单词要反转）
// 统计的结果是字典树的叶节点的高度+1 之和

func minimumLengthEncoding(words []string) int {
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(reverseStr(word))
	}
	//fmt.Println(trie)
	return trie.WalkAll()
}

type trieNode struct {
	value    rune
	children map[rune]*trieNode
	cnt      int
}

type Trie struct {
	root *trieNode
	size int
}

// NewTrie 创建 Trie 树
func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{children: make(map[rune]*trieNode)},
		size: 0,
	}
}

// NewTrieNode 创建一个 Trie 树的节点
func NewTrieNode(val rune) *trieNode {
	return &trieNode{
		value:    val,
		children: make(map[rune]*trieNode),
	}
}

func (t *Trie) Size() int {
	return t.size
}

// Insert方法： 给 Trie 树添加一个字符串
func (t *Trie) Insert(vals []rune) {
	cur := t.root
	t.size++
	for idx, val := range vals {
		if _, ok := cur.children[val]; !ok {
			cur.children[val] = NewTrieNode(val)
		}
		cur = cur.children[val]
		cur.cnt = idx + 1
	}
}

func Walk(p *trieNode) int {
	sum := 0
	if len(p.children) == 0 {
		sum = p.cnt + 1
	}
	for _, v := range p.children {
		sum += Walk(v)
	}
	return sum
}

func (t *Trie) WalkAll() int {
	return Walk(t.root)
}

// 翻转字符串
func reverseStr(s string) []rune {
	newS := []rune(s)
	sLen := len(s)

	for i := 0; i < sLen/2; i++ {
		newS[i], newS[sLen-i-1] = newS[sLen-i-1], newS[i]
	}
	return newS
}
