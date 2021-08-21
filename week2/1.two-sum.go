package main

func twoSum(nums []int, target int) []int {

	l := len(nums)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		num := target - nums[i]
		if _, ok := m[num]; ok { //如果另一个加数存在，被记录过，返回
			return []int{m[num], i}
		}
		m[nums[i]] = i // 将每一个值和他的下标，存起来，记录一下

	}
	return nil

}
