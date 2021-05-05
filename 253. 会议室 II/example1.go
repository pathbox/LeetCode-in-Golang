package LeetCode253

import (
	"math"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	meetingRooms := make(map[int]int)
	for _, interval := range intervals {
		meetingRooms[interval[0]] += 1
		meetingRooms[interval[1]] -= 1
	}

	times := make([]int, 0)
	for k := range meetingRooms {
		times = append(times, k)
	}
	sort.Ints(times)

	minRooms := 0
	currentRooms := 0
	for _, time := range times {
		currentRooms += meetingRooms[time]
		minRooms = int(math.Max(float64(minRooms), float64(currentRooms)))
	}
	return minRooms
}
