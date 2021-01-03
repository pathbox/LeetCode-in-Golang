package LeetCode303

type NumArray struct {
	presum []int // 存储 [0,i] 的和
}

// 一维前缀和
func Constructor(nums []int) NumArray {
	a := NumArray{}
	a.presum = append(a.presum, 0)
	for i := 1; i <= len(nums); i++ {
		t := a.presum[i-1] + nums[i-1]
		a.presum = append(a.presum, t)
	}
	return a

	// numa := NumArray{[]int{0}}
	// for i, v := range nums {
	// numa.Value = append(numa.Value, v + numa.Value[i])
	//}
	//return numa
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.presum[j+1] - this.presum[i]
}
