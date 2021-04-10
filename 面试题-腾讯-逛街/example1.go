package main

import "fmt"

/*
小Q在周末的时候和他的小伙伴来到大城市逛街，一条步行街上有很多高楼，共有n座高楼排成一行。
小Q从第一栋一直走到了最后一栋，小Q从来都没有见到这么多的楼，所以他想知道他在每栋楼的位置处能看到多少栋楼呢？（当前面的楼的高度大于等于后面的楼时，后面的楼将被挡住）
输入描述:
输入第一行将包含一个数字n，代表楼的栋数，接下来的一行将包含n个数字wi(1<=i<=n)，代表每一栋楼的高度。
1<=n<=100000;
1<=wi<=100000;
输出描述:
输出一行，包含空格分割的n个数字vi，分别代表小Q在第i栋楼时能看到的楼的数量。
输入例子1:
6
5 3 8 3 2 5
输出例子1:
3 3 5 4 4 4
例子说明1:
当小Q处于位置3时，他可以向前看到位置2,1处的楼，向后看到位置4,6处的楼，加上第3栋楼，共可看到5栋楼。当小Q处于位置4时，他可以向前看到位置3处的楼，向后看到位置5,6处的楼，加上第4栋楼，共可看到4栋楼。
*/

/*
原理其实是：从左到右遍历一个得到一个递增的单调栈，从右到左遍历得到一个递增的单调栈
https://www.liuyixiang.com/post/109198.html
*/
// func main() {
// 	arr := []int{5, 3, 8, 3, 2, 5}
// 	res := make([]int, 0)
// 	rightLook := make([]int, len(arr))
// 	stack := make([]int, 0)

// 	for i := len(arr) - 1; i >= 0; i-- { // 倒着遍历
// 		rightLook[i] = len(stack) // 第一个栈 从栈底到栈顶递减的栈，栈顶到栈底是从小到大，栈底是大数
// 		// 遍历到当前i，此时stack的元素个数就是以i向右看呈单调递增的楼的数量，也就是能看到的数量
// 		for len(stack) > 0 && arr[i] >= arr[stack[len(stack)-1]] {
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, i)
// 	}

// 	stack = make([]int, 0) // 重新置空 用于存第二个单调栈,也是从小到大，栈底是大数

// 	for i := 0; i < len(arr); i++ {
// 		total := rightLook[i] + 1 + len(stack)
// 		for len(stack) > 0 && arr[i] >= arr[stack[len(stack)-1]] { // 栈存的是要比当前i楼层还要大的且是当掉递增的单调栈
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, i)
// 		res = append(res, total)
// 	}

// 	fmt.Println(res)
// }

func main() {
	arr := []int{5, 3, 8, 3, 2, 5}
	res := make([]int, 0)

	rightLook := make([]int, len(arr))
	stack := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		rightLook[i] = len(stack)
		for len(stack) > 0 && arr[i] >= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		total := rightLook[i] + 1 + len(stack)
		for len(stack) > 0 && arr[i] >= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		res = append(res, total)
	}
	fmt.Println(res)
}
