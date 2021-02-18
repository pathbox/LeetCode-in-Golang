package LeetCode756

func pyramidTransition(bottom string, allowed []string) bool {
	// 用位图表示最底下这层
	bits := make([]byte, len(bottom))
	for i := range bottom {
		bits[i] |= 1 << (bottom[i] - 'A')
	}

	// 将allowed数组转换为二维数组
	matrix := make([][]byte, 8)
	for i := range matrix {
		matrix[i] = make([]byte, 8)
	}
	for i := range allowed {
		x := byte(allowed[i][0] - 'A')
		y := byte(allowed[i][1] - 'A')
		matrix[x][y] |= 1 << (allowed[i][2] - 'A')
	}
	return bfs(matrix, bits)
}

func bfs(matrix [][]byte, bits []byte) bool {
	size := len(bits)
	if size == 1 {
		return true
	}

	newBits := make([]byte, size-1) // 上一层的bits
	for i := range newBits {
		// 根据当前层第i和第i+1个节点的可能取值的组合 填充每一个父节点
		for bit := 0; bit < 8; bit++ {
			if bits[i]&(1<<bit) > 0 {
				for nextBit := 0; nextBit < 8; nextBit++ {
					if bits[i+1]&(1<<nextBit) > 0 {
						newBits[i] |= matrix[bit][nextBit]
					}
				}
			}
		}

		if newBits[i] == 0 { // 没有得到可能取值
			return false
		}
	}
	return bfs(matrix, newBits)
}
