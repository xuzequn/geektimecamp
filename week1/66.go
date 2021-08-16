package main

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- { // 低位到高位，是 从list的尾到头
		digits[i]++ // 当前位加一，初始加一 或者 进位加一

		if digits[i]%10 != 0 {
			return digits // 如果不等于10 ，返回digits
		} else {
			digits[i] = 0 // 如果等于10， 取0。
		}
	}
	result := []int{1} // 进位加一
	result = append(result, digits[0:]...)
	return result

}
