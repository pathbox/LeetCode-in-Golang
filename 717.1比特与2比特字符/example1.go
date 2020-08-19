package LeetCode717

// 三种字符分别为 0，10 和 11，那么 \mathrm{bits}bits 数组中出现的所有 0 都表示一个字符的结束位置（无论其为一比特还是两比特）。因此最后一位是否为一比特字符，只和他左侧出现的连续的 1 的个数（即它与倒数第二个 0 出现的位置之间的 1 的个数，如果 \mathrm{bits}bits 数组中只有 1 个 0，那么就是整个数组的长度减一）有关。如果 1 的个数为偶数个，那么最后一位是一比特字符，如果 1 的个数为奇数个，那么最后一位不是一比特字符

func isOneBitCharacter(bits []int) bool {
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			return (len(bits)-i)%2 == 0
		}
	}
	return (len(bits)-1)%2 == 0
}
