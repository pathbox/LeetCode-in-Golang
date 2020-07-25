package LeetCode1036

import (
	"container/list"
	"strconv"
)

// https://leetcode-cn.com/problems/escape-a-large-maze/solution/fei-chang-you-yi-si-de-bfs-by-zrd-4/

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) == 0 {
		return true
	}
	blocks := make(map[string]bool, len(blocked))
	for _, v := range blocked {
		if isSame(v, source) || isSame(v, target) {
			return false
		}
		blocks[pointToString(v)] = true
	}
	return bfs(source, target, blocks) && bfs(target, source, blocks)
}

func bfs(source, target []int, blocks map[string]bool) bool {
	const maxSize = 1e6
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := map[string]bool{pointToString(source): true}
	queue := list.New()
	queue.PushBack(source)
	for queue.Len() > 0 {
		if queue.Len() > len(blocks) {
			return true
		}
		point := queue.Remove(queue.Front()).([]int)
		for _, v := range dirs {
			r, c := point[0]+v[0], point[1]+v[1]
			if r < 0 || r >= maxSize || c < 0 || c >= maxSize {
				continue
			}
			nextPoint := []int{r, c}
			key := pointToString(nextPoint)
			if visited[key] || blocks[key] {
				continue
			}
			if isSame(nextPoint, target) {
				return true
			}
			visited[key] = true
			queue.PushBack(nextPoint)
		}
	}
	return false
}

func isSame(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func pointToString(point []int) string {
	return strconv.Itoa(point[0]) + strconv.Itoa(point[1])
}
