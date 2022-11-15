package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solution(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	v := TwoSum(nums, 9)
	assert.Equal(t, []int{0, 1}, v)

	nums = []int{3, 2, 4}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{1, 2}, v)

	nums = []int{2, 3, 3}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{1, 2}, v)

	nums = []int{3, 3}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{0, 1}, v)
}
