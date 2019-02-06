package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length = 15
	var list []int
	// 以时间戳为种子生成随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(100)))
	}

	fmt.Println("Before sort: ", list)
	result := ShellSort(list)
	fmt.Println("After sort: ", result)
}

func ShellSort(list []int) []int {
	// 这里就以n/2为增量
	gap := 2
	length := len(list)
	step := length / gap

	for step >= 1 {
		// 按步长开始每个分组的排序
		for i := step; i < length; i++ {
			// 将按步长分组的子队列用直接插入排序算法进行排序
			insertionSortByStep(list, step)
		}
		// 完成一轮后再次缩小增量
		step /= gap
	}
	return list
}

func insertionSortByStep(tree []int, step int) {
	for i := step; i < len(tree); i++ {
		for j := i; j >= step && tree[j] < tree[j-step]; j -= step {
			tree[j], tree[j-step] = tree[j-step], tree[j]
		}
	}
}
