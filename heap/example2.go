package main

import "fmt"

// 大顶堆， 父节点不小于左右子节点

type Heap struct {
	Size int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array
	return h
}

// 最大堆插入元素
func (h *Heap) Push(x int) {
	// 堆没有元素时，使元素成为顶点后退出
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	// i 是要插入节点的下标，一般从最尾部开始
	i := h.Size
	// 如果下标存在
	// 将小的值 x 一直上浮
	for i > 0 {
		parent := (i - 1) / 2 // i位置的父节点
		// 如果插入的值小于等于父亲节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x <= h.Array[parent] {
			break
		}
		// 否则将父亲节点与该节点互换，然后向上翻转，将最大的元素一直往上推
		h.Array[i] = h.Array[parent] // i的值为父节点的值
		i = parent                   // i的位置变为父节点
	}
	// 将该值 x 放在不会再翻转的位置
	h.Array[i] = x
	h.Size++
}

// 最大堆移除根节点元素，也就是最大的元素
func (h *Heap) Pop() int {
	// 没有元素，返回-1
	if h.Size == 0 {
		return -1
	}
	// 取出根节点
	ret := h.Array[0]

	// 因为根节点要被删除了，将最后一个节点放到根节点的位置上
	h.Size--
	x := h.Array[h.Size]
	h.Array[h.Size] = ret

	// 对根节点进行向下翻转，小的值 x 一直下沉，维持最大堆的特征
	i := 0
	for {
		// a，b为下标 i 左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 左儿子下标超出了，表示没有左子树，那么右子树也没有，直接返回
		if a >= h.Size {
			break
		}
		// 有右子树，拿到两个子节点中较大节点的下标
		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b
		}
		if x >= h.Array[a] {
			break // x到了合适的位置
		}

		// 将较大的儿子与父亲交换，维持这个最大堆的特征
		h.Array[i] = h.Array[a]
		i = a
	}

	h.Array[i] = x
	return ret
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

	// 构建最大堆
	h := NewHeap(list)
	for _, v := range list {
		h.Push(v)
	}

	// 将堆元素移除
	for range list {
		h.Pop()
	}

	// 打印排序后的值
	fmt.Println(list)
}

// 从堆构建到移除最坏和最好的时间复杂度：O(nlogn)
