package main

import (
	"fmt"
	"strconv"
	"strings"
)

const maxuint = ^uint32(0)
const maxint = int(maxuint >> 1)
const minint = -maxint - 1

func myAtoi(str string) int {

	str = strings.Trim(str, " ")
	if len(str) == 0 || !check(int32(str[0])) {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		j := 1
		for j < len(str) {
			c := int32(str[j])
			if c > 57 || c < 48 {
				break
			}
			j++
		}

		i, err = strconv.Atoi(str[:j])
		if err != nil {
			fmt.Println(err)
		}
	}

	if i > maxint {
		return maxint
	} else if i < minint {
		return minint
	} else {
		return i
	}

}

func check(s int32) bool {
	if s > 57 || s < 48 {
		if s != 43 && s != 45 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(myAtoi("-4193 with words"))
}
