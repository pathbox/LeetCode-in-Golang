package LeetCode381

import "math/rand"

type RandomizedCollection struct {
	rMap   map[int]map[int]int
	rSlice []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{make(map[int]map[int]int, 0), make([]int, 0)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (randCollection *RandomizedCollection) Insert(val int) bool {
	//更新slice
	randCollection.rSlice = append(randCollection.rSlice, val)
	if _, ok := randCollection.rMap[val]; !ok {
		//不存在，新建子map
		randCollection.rMap[val] = map[int]int{len(randCollection.rSlice) - 1: 1}
		return true
	} else {
		//存在，更新子map
		randCollection.rMap[val][len(randCollection.rSlice)-1] = 1
		return false
	}
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (randCollection *RandomizedCollection) Remove(val int) bool {
	//不存在
	if len(randCollection.rMap[val]) == 0 {
		return false
	}

	//随机获取一个val的下标
	index := 0
	for index, _ = range randCollection.rMap[val] {
		break
	}
	//把index与队尾交换，删除map中旧的index的对应关系，同时新增现在的index的对应关系
	randCollection.rSlice[index] = randCollection.rSlice[len(randCollection.rSlice)-1]
	//先删map，防止新增后又被删除，导致实际删除2次
	delete(randCollection.rMap[val], index)
	randCollection.rMap[randCollection.rSlice[index]][index] = 1

	//先从map中删除原队尾的对应关系，再删除slice队尾，先删map防止最后一个元素删除队尾后，队列为空导致数组越界panic
	delete(randCollection.rMap[randCollection.rSlice[index]], len(randCollection.rSlice)-1)
	randCollection.rSlice = randCollection.rSlice[:len(randCollection.rSlice)-1]

	return true
}

/** Get a random element from the collection. */
func (randCollection *RandomizedCollection) GetRandom() int {
	//rand.Intn参数小于等于0导致panic
	if len(randCollection.rSlice) == 0 {
		return -1
	}
	index := rand.Intn(len(randCollection.rSlice))
	return randCollection.rSlice[index]
}
