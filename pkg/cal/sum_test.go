package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1.執行 go test
// 測試目前目錄下的測試檔

// 2.執行 go test -v
// 查看測試函數名稱和運行時間

// 3.執行 go test ./…
// 測試包含子目錄下的測試檔

// 4.執行 go test -cover 或  go test ./... -cover
// 顯示覆蓋率

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
