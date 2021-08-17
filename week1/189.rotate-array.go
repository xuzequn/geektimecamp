package main

// 模拟
func rotate(nums []int, k int) {
	l := len(nums)
	temp := make([]int, l)
	for i, v := range nums {
		temp[(i+k)%l] = v
	}
	copy(nums, temp)

}
