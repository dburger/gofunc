package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	list := []int{1, 2, 3, 4}
	list = filter(list, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, list)
}

func TestTransform(t *testing.T) {
	list := []int{1, 2, 3, 4}
	list = transform(list, func(i int) int { return i * 2 })
	assert.Equal(t, []int{2, 4, 6, 8}, list)
}

func TestFlatmap(t *testing.T) {
	list := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5},
		[]int{6},
	}
	flattened := flatmap(list, func(i int) int {
		return i * 2
	})
	assert.Equal(t, []int{2, 4, 6, 8, 10, 12}, flattened)
}

func TestReduce(t *testing.T) {
	list := []int{1, 2, 3}
	result := reduce(list, func(a, b int) int { return a + b }, 0)
	assert.Equal(t, 6, result)

	result = reduce(list, func(a, b int) int { return a * b }, 1)
	assert.Equal(t, 6, result)
}
