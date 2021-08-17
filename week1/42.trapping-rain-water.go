package main

import "fmt"

// 暴力解法 时间复杂度 O(n*n)
// func trap(height []int) int {
// 	l := len(height)
// 	if l == 0 {
// 		return 0
// 	}
// 	result := 0
// 	front := 0
// 	rear := l - 1
// 	// 退出条件
// 	for front < rear {

// 		// 去掉首零
// 		for height[front] == 0 && front < l-1 {
// 			front++
// 		}
// 		// 去掉尾0
// 		for height[rear] == 0 && rear > 0 {
// 			rear--
// 		}
// 		for i := front; i <= rear; i++ {
// 			if height[i] == 0 { // 中间遇零加一
// 				result++
// 			} else { // 每轮减1
// 				height[i]--
// 			}
// 		}
// 	}
// 	return result
// }

// 单调栈的做法
func trap(height []int) int {
	stack := []int{} // 栈
	ans := 0
	for i, h := range height {
		// 当栈非空， 并且 当前元素比栈顶元素大，进行处理
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]   // 拿到栈顶元素
			stack = stack[:len(stack)-1] // 取出栈顶元素
			if len(stack) == 0 {         // 栈内元素为0退出， 相当于本次循环开始时栈内只有一个值，并且在这次循环中已经被取出，则退出
				break
			}
			left := stack[len(stack)-1]              // 拿到栈顶下面一个元素。
			w := i - left - 1                        // 求宽
			hh := min(h, height[left]) - height[top] // 求高，
			ans += w * hh
		}
		stack = append(stack, i) // 新元素入栈
	}
	return ans

}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func main() {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// height := []int{4, 2, 0, 3, 2, 5}
	height := []int{2, 0, 2}
	num := trap(height)
	fmt.Println(num)
}
