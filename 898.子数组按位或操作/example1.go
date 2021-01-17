package LeetCode898

func subarrayBitwiseORs(arr []int) int {
	set := make(map[int]int)
	for i, x := range arr {
		set[x] = 0
		for j := i - 1; j >= 0; j-- {
			if (x | arr[j]) == arr[j] {
				break
			}
			arr[j] |= x
			set[arr[j]] = 0
		}
	}
	return len(set)
}
