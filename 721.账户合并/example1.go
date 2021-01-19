package LeetCode721

import "sort"

type UnionFind struct {
	count  int
	parent []int
}

func (u *UnionFind) Init(count int) {
	u.count = count
	u.parent = make([]int, count)
	for i, _ := range u.parent {
		u.parent[i] = -1
	}
}

func (u *UnionFind) Find(node int) int {
	if u.parent[node] < 0 {
		return node
	}
	result := u.Find(u.parent[node])
	u.parent[node] = result
	return result
}

func (u *UnionFind) Union(node1 int, node2 int) (int, bool) {
	root1 := u.Find(node1)
	root2 := u.Find(node2)
	if root1 == root2 {
		return root2, false
	}
	u.parent[root1] = root2
	u.count--
	return root2, true
}

func accountsMerge(accounts [][]string) [][]string {
	var uf UnionFind
	uf.Init(len(accounts))
	emailToID := make(map[string]int, len(accounts))

	for i, account := range accounts {
		for _, val := range account[1:] {
			findIdx, ok := emailToID[val]
			if !ok {
				emailToID[val] = i
				continue
			}
			uf.Union(i, findIdx)
		}
	}

	tmpWork := make([][]string, len(accounts))
	for k, v := range emailToID {
		p := uf.Find(v)
		if tmpWork[p] == nil {
			tmpWork[p] = []string{k}
			continue
		}
		tmpWork[p] = append(tmpWork[p], k)
	}
	var result [][]string
	for i, r := range tmpWork {
		if r == nil {
			continue
		}
		sort.StringSlice(r).Sort()
		result = append(result, append([]string{accounts[i][0]}, r...))
	}

	return result
}
