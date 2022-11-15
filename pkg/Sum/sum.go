package sum

func TwoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, v := range nums {
		found, ok := m[target-v]
		if ok {
			return []int{found, i}
		}
		m[v] = i
	}

	return nil
}
