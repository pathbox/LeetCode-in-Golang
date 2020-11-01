package LeetCode875

import "math"

func canEat(piles []int, speed, h int) bool {
	sum := 0.00
	for _, v := range piles {
		sum += math.Ceil(float64(v) * 1.0 / float64(speed))
	}
	return sum > float64(h)
}
func minEatingSpeed(piles []int, H int) int {
	maxVal := 1
	for _, i := range piles {
		maxVal = int(math.Max(float64(maxVal), float64(i)))
	}
	low := 1
	height := maxVal
	for low < height {
		mid := (low + height) / 2
		if canEat(piles, mid, H) {
			low = mid + 1
		} else {
			height = mid
		}
	}
	return low // 可以有很多满足条件，找到满足的最左边界，最小的合适值
}

// // 这里直接套用模板代码
// func minEatingSpeed(piles []int, H int) int {
// 	left, right := 0, MaxOf(piles)
// 	for left < right {
// 		mid := left + (right-left)>>1
// 		if TotalTime(piles, mid) <= H { // mid 的含义是「吃香蕉的速度」
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return left // 可以有很多满足条件，找到满足的最左边界，最小的合适值
// }

// func TotalTime(piles []int, k int) int {
// 	time := 0
// 	for _, v := range piles {
// 		time += v / k
// 		if v%k > 0 {
// 			time++
// 		}
// 	}
// 	return time
// }

// func MaxOf(nums []int) int {
// 	m := -1 << 63
// 	for _, v := range nums {
// 		m = Max(m, v)
// 	}
// 	return m
// }

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
