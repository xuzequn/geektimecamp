package main

// 哈希

func twoSum(nums []int, target int) []int {
	indexlist := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v := target - nums[i]

		if _, ok := indexlist[v]; ok {
			return []int{indexlist[v], i}
		}
		// 哈希的key 是 nums的元素
		indexlist[nums[i]] = i
	}
	return nil
}
