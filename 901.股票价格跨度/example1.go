package LeetCode901

// 单调递减栈

type StockSpanner struct {
	weight []int
	price  []int // 单调递减栈
}

func Constructor() StockSpanner {
	return StockSpanner{
		weight: make([]int, 0),
		price:  make([]int, 0),
	}
}

func (ss *StockSpanner) Next(price int) int {
	w := 1
	for len(ss.price) > 0 && ss.price[len(ss.price)-1] <= price {
		ss.price = ss.price[:len(ss.price)-1] // 小于price的全部出栈
		w += ss.weight[len(ss.weight)-1]
		ss.weight = ss.weight[:len(ss.weight)-1]
	}

	ss.price = append(ss.price, price)
	ss.weight = append(ss.weight, w)
	return w
}
