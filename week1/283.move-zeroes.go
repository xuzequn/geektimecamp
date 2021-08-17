package main

func moveZeroes(nums []int) {
	a := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[a] = nums[a], nums[i]
			a++
		}
	}

}
