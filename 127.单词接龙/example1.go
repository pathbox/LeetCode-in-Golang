package LeetCode127

func ladderLength(beginWord string, endWord string, wordList []string) int {
	step := 0
	startQueue := []string{beginWord}
	endQueue := []string{endWord}
	startUsed := make([]bool, len(wordList))
	endUsed := make([]bool, len(wordList))
	wordMap := make(map[string]int, 0)
	for i, w := range wordList {
		wordMap[w] = i
	}
	if idx, ok := wordMap[endWord]; !ok {
		return 0
	} else {
		endUsed[idx] = true
	}

	for len(startQueue) > 0 {
		step++
		l := len(startQueue)
		for i := 0; i < l; i++ {
			chars := []byte(startQueue[i])
			for j := 0; j < len(chars); j++ {
				o := chars[j]
				for c := 'a'; c <= 'z'; c++ {
					chars[j] = byte(c)
					idx, ok := wordMap[string(chars)]
					if !ok || startUsed[idx] {
						continue
					} else {
						if endUsed[idx] {
							return step + 1
						}
						startQueue = append(startQueue, wordList[idx])
						startUsed[idx] = true
					}
				}
				chars[j] = o
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			startUsed, endUsed = endUsed, startUsed
		}
	}
	return 0
}
