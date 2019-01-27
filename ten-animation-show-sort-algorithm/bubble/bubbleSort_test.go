package bubble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	ary := []int{5, 2, 3, 8, 4}
	ast := assert.New(t)

	result := BubbleSort(ary)
	ast.Equal(2, result[0], "输入:%v", ary)
}
