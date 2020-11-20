package basic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"magical.dev/tesing/basic"
)

func TestSimplePlus(t *testing.T) {
	result := basic.Plus(1, 2)

	assert.Equal(t, 3, result)
}
