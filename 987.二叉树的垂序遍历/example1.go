package LeetCode987

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	result := [][]int{}
	posMap := getNodePosition(root)
	sli := []int{}
	for key, _ := range posMap {
		sli = append(sli, key)
	}
	sort.Ints(sli)
	for _, v := range sli {
		tmp := posMap[v]
		result = append(result, tmp)
	}
	return result
}

func getNodePosition(root *TreeNode) map[int][]int {
	positionMap := make(map[int][]int)
	nodeList := []*TreeNode{root}
	yList := []int{0}
	for len(nodeList) > 0 {
		l := len(nodeList)
		mapTmp := make(map[int][]int)
		for i := 0; i < l; i++ {
			nodeTmp := nodeList[0]
			yTmp := yList[0]
			mapTmp[yTmp] = append(mapTmp[yTmp], nodeTmp.Val)
			if nodeTmp.Left != nil {
				nodeList = append(nodeList, nodeTmp.Left)
				yList = append(yList, yTmp-1)
			}
			if nodeTmp.Right != nil {
				nodeList = append(nodeList, nodeTmp.Right)
				yList = append(yList, yTmp+1)
			}
			nodeList = nodeList[1:]
			yList = yList[1:]
		}
		for key, value := range mapTmp {
			if len(value) == 1 {
				positionMap[key] = append(positionMap[key], value[0])
			} else {
				sort.Ints(value)
				positionMap[key] = append(positionMap[key], value...)
			}
		}
	}
	return positionMap
}
