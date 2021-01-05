package LeetCode218

import "sort"

type boxing struct {
	node, pos, height, side int
}

func getSkyline(buildings [][]int) [][]int {
	// 标记并排序输入形状所有的拐点，每个矩形两个点。
	var box []boxing
	for i, v := range buildings {
		box = append(box, boxing{i, v[0], v[2], 0}, boxing{i, v[1], v[2], 1})
	}

	sort.Slice(box, func(i, j int) bool {
		return box[i].pos < box[j].pos // 左边节点会在后边节点前面
	})

	var hill, result [][]int

	maxhill := -1
	for i := 0; i < len(box); i++ {
		v := box[i]
		if v.side == 0 { // 左边节点 入队列
			// push 拐点
			hill = append(hill, []int{v.height, v.node})
		} else {
			// pop 拐点
			for id, vv := range hill {
				if vv[1] == v.node {
					hill = append(hill[:id], hill[id+1:]...)
					break
				}
			}
		}
		// 当前位置（x)，最高点(y)查找。
		maxh := 0
		for _, vv := range hill {
			if vv[0] > maxh {
				maxh = vv[0]
			}
		}

		if !(i < len(box)-1 && box[i].pos == box[i+1].pos) { // 同一个位置（x）的出现多个（Y）的处理
			// 当前最高点和历史最高点（前序）存在变化Delta则必须记录之（Result）
			if maxhill < maxh || maxhill > maxh {
				result = append(result, []int{v.pos, maxh})
			}
			maxhill = maxh
		}
	}
	return result
}
