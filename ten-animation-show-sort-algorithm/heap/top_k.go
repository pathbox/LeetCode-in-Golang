package topk

func sift(li []int, low, high int) {
	i := low
	j := 2*i + 1
	tmp := li[i]
	for j <= high {
		if j < high && li[j] > li[j+1] {
			j++
		}
		if tmp > li[j] {
			li[i] = li[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	li[i] = tmp
}

func top_k(li []int, k int) []int {
	for i := k/2 - 1; i >= 0; i-- {
		sift(li, i, k-1)
	}
	for j := k; j < len(li); j++ {
		if li[0] < li[j] {
			li[0], li[j] = li[j], li[0]
			sift(li, 0, k-1)
		}
	}
	for n := k - 1; n > 0; n-- {
		li[0], li[n] = li[n], li[0]
		sift(li, 0, n-1)
	}
	return li[:k]
}

/*
用堆排解决top_k问题(优先队列)，思路：

　　a.先取前k个数建小根堆，这样就能保证堆顶元素是整个堆的最小值;

　　b.然后遍历列表的k到最后，如果值比堆顶大，就和堆顶交换，交换完后再对堆建小根堆，这样就能保证交换完后，堆顶元素仍然是整个堆的最小值；

　　c.一直遍历到列表的最后一个值，这一步做完之后，就保证了整个列表最大的前k个数已经放进了堆里；

　　d.按数的大到小出堆；
*/
