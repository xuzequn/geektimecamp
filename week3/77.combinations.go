package main

import "fmt"

func combine(n int, k int) [][]int {

	result := [][]int{}
	visited := make([]bool, n+1)
	fmt.Println(visited)
	var dfs func(i, deep int)
	path := []int{}
	dfs = func(i, deep int) {
		if !visited[i] {
			visited[i] = true
			path = append(path, i)
		}
		// 终止条件
		if deep == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return

		}
		// 找到下一个路径节点
		for j := i + 1; j <= n; j++ {
			dfs(j, deep+1)
			visited[j] = false
			path = path[:len(path)-1]
		}
	}
	// 因为组合不考虑顺序，不能重复，所以第一个数只能处理一次
	for i := 1; i <= n; i++ {
		dfs(i, 1)
		path = path[:len(path)-1]
	}
	return result

}
