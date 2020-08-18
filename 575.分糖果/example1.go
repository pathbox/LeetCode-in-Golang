package LeetCode575

/*
这题整体思路比较简单，如果要妹妹尽量多的拿到不同种类的糖果，那么，只要看到不同种类，直接给妹妹就行了。注意一点，如果现在妹妹拥有的糖果已经是整体的一半了，那么直接返回就好了。
哈希表的存在主要是为了判别当前种类的糖果是否应该分配给妹妹。如果已经拿过了，就跳到下一个。
*/

func distributeCandies(candies []int) int {
	have := make(map[int]bool)
	types := 0
	for i := 0; i < len(candies); i++ {
		if types >= len(candies)/2 {
			break
		}
		if _, ok := have[candies[i]]; !ok {
			types++
			have[candies[i]] = true
		}
	}
	return types
}
