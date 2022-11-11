package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	val := 5
	assert.Equal(t, 5, val)
}
