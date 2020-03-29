package LeetCode820

func minimumLengthEncoding(words []string) int {
	var (
		root = NewTree()
	)
	for i := 0; i < len(words); i++ {
		temp := words[i]
		insert(root, temp)
	}
	return findLeaf(root, 0)
}

type node struct {
	isStr bool
	next  []*node
}

func NewTree() *node {
	return &node{
		isStr: false,
		next:  make([]*node, 26), // 只能用于存小写字母,每个索引位置标示对应的小写字母
	}
}

func insert(root *node, str string) {
	for i := len(str) - 1; i >= 0; i-- {
		if root.next[str[i]-'a'] == nil {
			root.next[str[i]-'a'] = &node{
				isStr: i == 0,
				next:  make([]*node, 26),
			}
		}
		root = root.next[str[i]-'a']
	}
}

func findLeaf(root *node, length int) int {
	var (
		isLeaf = true
		sum    = 0
	)
	nodes := root.next
	for i := 0; i < 26; i++ {
		if nodes[i] != nil {
			isLeaf = false
			sum += findLeaf(nodes[i], length+1)
		}
	}
	if isLeaf {
		return length + 1
	}
	return sum
}

// 使用字典树解决后缀重复（单词要反转）
// 统计的结果是字典树的叶节点的高度+1 之和
