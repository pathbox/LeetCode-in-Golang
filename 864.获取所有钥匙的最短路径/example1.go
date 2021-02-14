package LeetCode864

import "container/list"

type Node struct {
	i, j int
	step int
	key  int
}

var (
	rand [][]int = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
)

func shortestPathAllKeys(grid []string) int {
	if len(grid) == 0 && len(grid[0]) == 0 {
		return -1
	}
	g := make([][]byte, len(grid))
	l := list.New()
	success := 0
	for i := 0; i < len(grid); i++ {
		g[i] = make([]byte, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			g[i][j] = grid[i][j]
			if g[i][j] >= 'a' && g[i][j] <= 'z' {
				success |= 1 << uint(g[i][j]-'a')
			}
			if g[i][j] == '@' {
				l.PushBack(Node{i, j, 0, 0})
			}
		}
	}
	visited := make([][][]bool, len(g)) // i, j, keyState
	for i := 0; i < len(g); i++ {
		visited[i] = make([][]bool, len(g[0]))
		for j := 0; j < len(g[0]); j++ {
			visited[i][j] = make([]bool, success+1)
		}
	}
	return bfs(success, g, l, visited)
}

func bfs(success int, g [][]byte, l *list.List, visited [][][]bool) int {
	if l.Len() <= 0 {
		return -1
	}
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		v := e.Value.(Node)
		i := v.i
		j := v.j

		if g[i][j] >= 'a' && g[i][j] <= 'z' {
			v.key |= 1 << uint(g[i][j]-'a')
		}
		if v.key == success {
			return v.step
		}
		for x := 0; x < len(rand); x++ {
			ii := i + rand[x][0]
			jj := j + rand[x][1]
			if ii < 0 || jj < 0 || ii >= len(visited) || jj >= len(visited[0]) {
				continue
			}
			if g[ii][jj] == '#' || visited[ii][jj][v.key] {
				continue
			}
			if g[ii][jj] >= 'A' && g[ii][jj] <= 'Z' {
				needKey := g[ii][jj] - 'A'
				if ((v.key >> needKey) & 1) == 0 {
					continue
				}
			}
			visited[ii][jj][v.key] = true
			n := Node{ii, jj, v.step + 1, v.key}
			l.PushBack(n)
		}
	}
	return -1
}
