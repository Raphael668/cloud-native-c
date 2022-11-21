package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bigger10Add(t *testing.T) {
	sum, err := Bigger10Add(11, 11)
	assert.Nil(t, err)
	assert.Equal(t, 22, sum)

	sum, err = Bigger10Add(55, 100)
	assert.Nil(t, err)
	assert.Equal(t, 155, sum)

	_, err = Bigger10Add(-1, 100)
	assert.NotNil(t, err)

	_, err = Bigger10Add(2, 3)
	assert.NotNil(t, err)

	_, err = Bigger10Add(5, 11)
	assert.NotNil(t, err)

	_, err = Bigger10Add(10, 10)
	assert.NotNil(t, err)
}
