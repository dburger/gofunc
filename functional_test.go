package gofunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	items := []int{1, 2, 3, 4}
	items = Filter(items, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, items)
}

func TestMap(t *testing.T) {
	items := []int{1, 2, 3, 4}
	items = Map(items, func(i int) int { return i * 2 })
	assert.Equal(t, []int{2, 4, 6, 8}, items)
}

func TestFlatmap(t *testing.T) {
	items := [][]int{
		{1, 2, 3},
		{4, 5},
		{6},
	}
	flattened := FlatMap(items, func(i int) int {
		return i * 2
	})
	assert.Equal(t, []int{2, 4, 6, 8, 10, 12}, flattened)
}

func TestReduce(t *testing.T) {
	items := []int{1, 2, 3}
	result := Reduce(items, func(a, b int) int { return a + b }, 0)
	assert.Equal(t, 6, result)

	result = Reduce(items, func(a, b int) int { return a * b }, 1)
	assert.Equal(t, 6, result)
}
