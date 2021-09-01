package main

import "strconv"

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}
	// 记录得到过的结果
	visitedpath := map[string]bool{}
	// 创建搜索函数
	var dfs func(res []int)
	dfs = func(res []int) {
		//找到一条路径 返回结果
		if len(nums) == len(res) {
			temp := make([]int, len(res)) // []int 初始化
			copy(temp, res)
			// 比46题 加一个结果判断重复，
			var str string
			// 先将结果专程一个字符串
			for _, v := range temp {
				str = str + strconv.Itoa(v)
			}
			// 判断是否得到过这个结果，如果没有添加，并记录这个结果
			_, ok := visitedpath[str]
			if !ok {
				result = append(result, res)
				visitedpath[str] = true
			}
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
