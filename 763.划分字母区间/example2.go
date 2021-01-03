package LeetCode763

// https://leetcode-cn.com/problems/partition-labels/solution/hua-fen-zi-mu-qu-jian-by-leetcode-solution/
func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0

	for i, c := range s {
		if lastPos[c-'a'] > end { // it makes end is max last position
			end = lastPos[c-'a']
		}
		// current index is equal to max last position, do partition now 扫到这个位置就切割，切出的字符不会在之后出现。
		if i == end { // it means it goes to one element last position here, so just pertition here
			partition = append(partition, end-start+1) // you got one partition, and goes to find next one
			start = end + 1
		}
	}
	return
}

// 时间复杂度：O(n)
// 空间复杂度：O(1) 使用的hash数组是固定大小
