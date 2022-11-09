package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PositiveAdd(t *testing.T) {
	sum, err := PositiveAdd(0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, sum)

	sum, err = PositiveAdd(1, 100)
	assert.Nil(t, err)
	assert.Equal(t, 101, sum)

	_, err = PositiveAdd(-1, 100)
	assert.NotNil(t, err)
}
