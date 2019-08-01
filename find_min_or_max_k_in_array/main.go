package main
import "fmt"

func patition(ary []int, left,right int) int {
	x := ary[left]
	for left < right {
		for left < right && ary[right] > x {
			right = right-1
		}
		for left < right && ary[left] < x {
			left = left + 1
		}
		ary[left], ary[right] = ary[right], ary[left]
	}
	ary[left] = x // 如果是找第k大的元素， 则 ary[right] = x return right
	return left
}

// 在left - right 中寻找
func find(ary []int, left, right, k int) {
	pos := patition(ary, left, right)
	if pos == k {
		fmt.Printf("%d\n", ary[pos])
	} else if pos > k {
		find(ary, left, pos-1, k)
	} else {
		find(ary, pos+1, right, k)
	}
}

func main() {
	k := 3
	ary := []int{5,3,7,20,25,1,11}
	find(ary,0,len(ary)-1,k)
}

/*
为什么算法时间是O(n)，因为这样的分治方式，遍历的次数差不多就是数组的长度n，
最多只要将整个数组遍历一次就能找到结果，而排序的方式至少都要 nlog(n)
*/