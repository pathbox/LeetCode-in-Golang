167. Two Sum II - Input array is sorted

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

You may assume that each input would have exactly one solution and you may not use the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2

1. Two pointer 方法

由于 给定的array数组是已经按从小到达排序了的（从达到小排序也可以，思路是类似的），有了顺序，这是一个必要条件

```go
l, r := 0, len(nums)-1 // 获得头尾两端的数组索引
for l < r { // 不断进行循环
  s := nums[l] + nums[r]
  if s == target {  // 相等，说明找到了
    return []int{l + 1, r + 1}
  } else if s < target { // 如果和小于target，说明和不够，应该头部往后走，以达到增大和的目的
    l++
  } else { // 如果和大于target，说明和太大，应该尾部往前走，以达到减小和的目的
    r--
  }
}
```

2. 二分搜索

为什么会联想到使用二分搜索的方式，因为数组是有序的。这里使用二分搜索，对于有重复元素的情况，会找到重复的元素值。比如：

nums := []int{1, 2, 3, 4, 4, 9, 56, 90}
target := 8

所以，有了一个小tip进行了弥补

3. 关联 [1.Two Sum](./1.two-sum)