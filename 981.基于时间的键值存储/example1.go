package LeetCode981

type TimeMap struct {
	m map[string][]node
}

type node struct {
	t   int
	val string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{make(map[string][]node)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], node{
		t:   timestamp,
		val: value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	nodes := this.m[key]
	l, r := 0, len(nodes)

	for l < r {
		mid := (l + r) >> 1
		if nodes[mid].t < timestamp {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l >= 0 && l < len(nodes) && nodes[l].t == timestamp {
		return nodes[l].val
	}
	if l > 0 {
		return nodes[l-1].val
	}
	return ""
}
