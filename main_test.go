package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BiggerAdd(t *testing.T) {
	th := 10
	sum, err := BiggerAdd(th, 11, 11)
	assert.Nil(t, err)
	assert.Equal(t, 22, sum)

	sum, err = BiggerAdd(th, 55, 100)
	assert.Nil(t, err)
	assert.Equal(t, 155, sum)

	_, err = BiggerAdd(th, -1, 100)
	assert.NotNil(t, err)

	_, err = BiggerAdd(th, 2, 3)
	assert.NotNil(t, err)

	_, err = BiggerAdd(th, 10, 10)
	assert.NotNil(t, err)

	// .. //
	th = -1
	sum, err = BiggerAdd(th, 11, 11)
	assert.Nil(t, err)
	assert.Equal(t, 22, sum)

	sum, err = BiggerAdd(th, 55, 100)
	assert.Nil(t, err)
	assert.Equal(t, 155, sum)

	_, err = BiggerAdd(th, -1, 100)
	assert.NotNil(t, err)

	_, err = BiggerAdd(th, -1, -1)
	assert.NotNil(t, err)

	_, err = BiggerAdd(th, -2, -6)
	assert.NotNil(t, err)
}
