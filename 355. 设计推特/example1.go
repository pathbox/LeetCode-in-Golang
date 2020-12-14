package LeetCode355

import "sort"

type Twitter struct {
	Time    int                 //全局时间戳
	Follows map[int]map[int]int //每个用户的关注池，3个int分别为：主用户id、粉丝的用户 id、被关注的用户是否存在（冗余字段，可以不管）
	Tweets  map[int]map[int]int //每个用户的tweet池, 3个int分别为：主用户id、tweet id、 tweet 时间戳
}

type Node struct {
	Tid   int
	Ttime int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Time:    0,
		Follows: make(map[int]map[int]int, 0),
		Tweets:  make(map[int]map[int]int, 0),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	t, ok := this.Tweets[userId] //获取该用户的 tweet 池
	if !ok {
		t = make(map[int]int, 0)
		this.Tweets[userId] = t
	}
	this.Time++
	t[tweetId] = this.Time //发tweet
}

/** userId follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(userId int, followeeId int) {
	if userId == followeeId { // 直接过滤掉“自己关注自己”这种情况（对输出结果无影响），简单粗暴
		return
	}
	f, ok := this.Follows[userId]
	if !ok {
		f = make(map[int]int, 0)
		this.Follows[userId] = f
	}
	f[followeeId] = 1 // 关注
}

/** userId unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(userId int, followeeId int) {
	if f, ok := this.Follows[userId]; ok { // 获取该用户的关注池
		delete(f, followeeId) // 取消关注
	}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	ids := []int{userId}
	if f, ok := this.Follows[userId]; ok { //获取该用户的关注池
		for followID, _ := range f {
			ids = append(ids, followID)
		}
	}
	//汇总所有用户的 tweet
	nodes := make([]*Node, 0)
	for _, id := range ids { //遍历所有用户
		if t, ok := this.Tweets[id]; ok { //获取该用户的关注池
			for tid, ttime := range t {
				nodes = append(nodes, &Node{tid, ttime})
			}
		}
	}
	//取最新发布的10个 tweet。 排序部分可以改用堆来进一步优化性能，不过懒得写了。
	ret := []int{}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Ttime > nodes[j].Ttime
	})
	for i := 0; i < len(nodes) && i < 10; i++ {
		ret = append(ret, nodes[i].Tid)
	}
	return ret
}
