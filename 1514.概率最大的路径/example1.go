package LeetCode1514

type Edge struct {
	val  int
	prob float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    nums:=make([][]Edge,n)
    for i,edge:=range edges{
    	from,to:=edge[0],edge[1]
    	if nums[from]==nil{
    		nums[from] = []Edge{{val: to,prob: succProb[i]}}
			}else {
				nums[from] = append(nums[from],Edge{val: to,prob: succProb[i]})
			}
			if nums[to]==nil{
				nums[to] =[]Edge{{val: from,prob: succProb[i]}}
			}else {
				nums[to] = append(nums[to],Edge{val: from,prob: succProb[i]})
			}
		}
	prob:=make([]float64,n)
	prob[start] = 1.0
	nodes:=make([]Edge,0)
	nodes = append(nodes,Edge{val: start,prob: 1.0})
	for len(nodes)>0{
		node:=nodes[0]
		nodes=nodes[1:]
		if prob[node.val]>node.prob{
			continue
		}
		for _,num:=range nums[node.val]{
			if prob[num.val]<prob[node.val]*num.prob{
				prob[num.val] = prob[node.val]*num.prob
				nodes = append(nodes,Edge{num.val,prob[num.val]})
			}
		}
	}
	return prob[end]
}
