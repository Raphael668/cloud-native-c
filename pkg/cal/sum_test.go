package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TwoSum(t *testing.T) {

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

	nums = []int{3}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{}, v)

	nums = []int{6}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{}, v)

	nums = []int{7, 1, 9, 11}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{}, v)

	nums = []int{}
	v = TwoSum(nums, 6)
	assert.Equal(t, []int{}, v)
}
