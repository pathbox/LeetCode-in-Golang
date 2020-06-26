package Offer045

import (
	"sort"
	"strconv"
)

type Slice []int

func (p Slice) Len() int { return len(p) }

func (p Slice) Less(i, j int) bool {
	istr := strconv.Itoa(p[i])
	jstr := strconv.Itoa(p[j])

	s1 := istr + jstr
	s2 := jstr + istr

	if s1 < s2 { // 转成字符串后进行比较
		return true
	}

	return false
}

func (p Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func minNumber(nums []int) string {
	var nums2 Slice = Slice(nums)
	sort.Sort(nums2)
	var ret string
	for _, num := range nums2 {
		ret += strconv.Itoa(num)
	}
	return ret
}
