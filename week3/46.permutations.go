package main

func permute(nums []int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}
	// 创建搜索函数
	var dfs func(res []int)
	dfs = func(res []int) {
		//找到一条路径 返回结果
		if len(nums) == len(res) {
			temp := make([]int, len(res)) // []int 初始化
			copy(temp, res)
			result = append(result, temp)
			return
		}
		// 将元素加入路径
		for i := 0; i < len(nums); i++ {
			// 先判断是否访问过
			if visited[i] == true {
				continue
			}
			res = append(res, nums[i])
			visited[i] = true
			dfs(res)
			res = res[:len(res)-1]
			visited[i] = false
		}
		return

	}

	// 开始搜索
	dfs([]int{})
	// 返回结果

	return result

}
