package LeetCode1306

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	queue := []int{start}

	for len(queue) > 0 {
		length := len(queue)
		for cur := 0; cur < length; cur++ {
			i := queue[cur]
			if !visited[i] {
				visited[i] = true
				if i+arr[i] < len(arr) {
					if arr[i+arr[i]] == 0 {
						return true
					}
					queue = append(queue, i+arr[i])
				}
				if i-arr[i] >= 0 {
					if arr[i-arr[i]] == 0 {
						return true
					}
					queue = append(queue, i-arr[i])
				}
			}
		}
		queue = queue[length:] // 出队列，进行下一轮遍历
	}
	return false
}
