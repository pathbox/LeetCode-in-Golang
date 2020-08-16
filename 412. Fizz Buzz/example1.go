package LeetCode412

import "strconv"

func fizzBuzz(n int) []string {
	var rst []string
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			rst = append(rst, "FizzBuzz")
		} else if i%5 == 0 {
			rst = append(rst, "Buzz")
		} else if i%3 == 0 {
			rst = append(rst, "Fizz")
		} else {
			rst = append(rst, strconv.Itoa(i))
		}
	}
	return rst
}
