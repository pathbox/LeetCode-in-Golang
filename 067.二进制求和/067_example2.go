package LeetCode067

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	x, _ := new(big.Int).SetString(a, 2)
	y, _ := new(big.Int).SetString(b, 2)

	zero, _ := new(big.Int).SetString("0", 2)

	for y.Cmp(zero) != 0 {
		ans := new(big.Int).Xor(x, y)
		car := x.And(x, y).Lsh(x, 1)
		x, y = ans, car
	}
	return fmt.Sprintf("%b", x)
}
