package LeetCode528

import (
	"math/rand"
	"time"
)

type Solution struct {
	PreSum []int
	Sum    int
	size   int
}

func Constructor(w []int) Solution {
	// 构造前缀和数组
	node := make([]int, len(w)+1)
	for i, c := range w {
		node[i+1] = c + node[i]
	}
	return Solution{
		PreSum: node[1:],
		Sum:    node[len(w)],
		size:   len(w),
	}
}

func (this *Solution) PickIndex() int {
	// 取[1, thie.Sum]的数:
	rand.Seed(time.Now().UnixNano())
	tmp := rand.Intn(this.Sum-1+1) + 1
	left := 0
	right := this.size
	w := this.PreSum
	// 二分法找左边界
	for left+1 < right {
		mid := left + (right-left)>>1
		if w[mid] >= tmp {
			right = mid
		} else {
			left = mid
		}
	}
	if w[left] >= tmp {
		return left
	} else {
		return right
	}
}
