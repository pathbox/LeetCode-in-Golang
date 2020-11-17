package LeetCode1636

import "sort"

type Num struct {
	num  int
	freq int
}

func frequencySort(nums []int) []int {
	//1. 统计每个数字出现的频数
	numsMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num]++
		} else {
			numsMap[num] = 1
		}
	}

	numsFreq := make([]Num, 0)
	for key, val := range numsMap {
		numsFreq = append(numsFreq, Num{key, val})
	}

	sort.Slice(numsFreq, func(i, j int) bool {
		if numsFreq[i].freq == numsFreq[j].freq {
			return numsFreq[i].num > numsFreq[j].num
		}
		return numsFreq[i].freq < numsFreq[j].freq
	})
	res := make([]int, 0)
	for _, val := range numsFreq {
		for i := 0; i < val.freq; i++ {
			res = append(res, val.num)
		}
	}
	return res
}
