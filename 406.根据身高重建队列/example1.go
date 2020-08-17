package LeetCode406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// 排序，先按照身高h排序，如果身高一样高，那就按照人数k排序
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0]) || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	// 按照k值插入到index=k的地方，index之后的往后移动
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}
	return people
}
