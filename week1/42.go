package main

import "fmt"

// 暴力解法 时间复杂度 O(n*n)
func trap(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	result := 0
	front := 0
	rear := l - 1
	// 退出条件
	for front < rear {

		// 去掉首零
		for height[front] == 0 && front < l-1 {
			front++
		}
		// 去掉尾0
		for height[rear] == 0 && rear > 0 {
			rear--
		}
		for i := front; i <= rear; i++ {
			if height[i] == 0 { // 中间遇零加一
				result++
			} else { // 每轮减1
				height[i]--
			}
		}
	}
	return result
}

func main() {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// height := []int{4, 2, 0, 3, 2, 5}
	height := []int{2, 0, 2}
	num := trap(height)
	fmt.Println(num)
}
