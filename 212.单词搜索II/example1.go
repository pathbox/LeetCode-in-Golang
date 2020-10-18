package LeetCode212

type TrieNode struct {
	word     string
	children [26]*TrieNode // 用map也可以
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	// 构造Trie树
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{} // 不存在则新建一个节点
			}
			node = node.children[c-'a'] // node会遍历到为是w单词最后一个字符
		}
		node.word = w // node会遍历到为是w单词最后一个字符
	}

	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, root, &result) // 开始dfs root就是trie树
		}
	}
	return result
}

func dfs(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) { // 超出边界的条件
		return
	}
	c := board[i][j]                             // 当前字符
	if c == '#' || node.children[c-'a'] == nil { // 访问过了或者字典中没有
		return
	}
	node = node.children[c-'a'] // 遍历trie树 取当前node的trie树节点

	if node.word != "" { // 说明在trie树中找到了匹配的单词
		*result = append(*result, node.word)
		node.word = "" // 防止重复添加
	}

	board[i][j] = '#'                // 在当前递归层访问过了
	dfs(i+1, j, board, node, result) // 以当前node为root继续寻找遍历
	dfs(i-1, j, board, node, result)
	dfs(i, j+1, board, node, result)
	dfs(i, j-1, board, node, result)
	board[i][j] = c // 回溯
}

/*
func dfs(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	// 1. 超出边界的条件的过滤
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]                             // 当前字符
	if c == '#' || node.children[c-'a'] == nil { // 访问过了或者字典中没有
		return
	}

	// 2. 具体操作逻辑，包括满足条件的操作

	board[i][j] = '#'                // 在当前递归层访问过了

	// 3. 四个方向的继续dfs
	dfs(i+1, j, board, node, result)
	dfs(i-1, j, board, node, result)
	dfs(i, j+1, board, node, result)
	dfs(i, j-1, board, node, result)

	// 4. 回溯
	board[i][j] = c // 回溯
}

*/
