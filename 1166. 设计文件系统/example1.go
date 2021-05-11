package LeetCode1166

import "strings"

type FileSystem struct {
	next  map[string]*FileSystem
	value int
}

func Constructor() FileSystem {
	return FileSystem{
		next: make(map[string]*FileSystem),
	}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if len(path) < 2 {
		return false
	}
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++ { //注意从1开始，第一个是空字符串
		if _, ok := this.next[paths[i]]; !ok {
			if i == len(paths)-1 { // 遍历到最后一个元素,将最后一个元素加入到字典树上
				this.next[paths[i]] = &FileSystem{
					next:  make(map[string]*FileSystem),
					value: value,
				}
				return true
			} else {
				return false // 说明没能遍历到最后一个元素,之前的字典树不匹配
			}
		} else {
			this = this.next[paths[i]] // 往下遍历next
		}
	}
	return false
}

func (this *FileSystem) Get(path string) int {
	paths := strings.Split(path, "/")

	for i := 1; i < len(paths); i++ {
		if _, ok := this.next[paths[i]]; ok {
			this = this.next[paths[i]]
		} else { // 说明字典树中间就不匹配了
			return -1
		}
	}
	return this.value
}
