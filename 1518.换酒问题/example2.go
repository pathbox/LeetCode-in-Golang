package LeetCode1518

func numWaterBottles(numBottles int, numExchange int) int {
	//一共喝的瓶数
	m := numBottles
	//这里循环一次代表喝了一瓶
	//这里的numBottles代表的是空瓶子，
	for numBottles >= numExchange {
		//喝一次m加一
		m++
		//喝了一瓶代表空瓶数减去numExchange，并且这个时候因为喝了一瓶所以会增加一个空瓶子所以还需要给空瓶加一
		numBottles = numBottles - numExchange + 1

	}
	return m
}

func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}
