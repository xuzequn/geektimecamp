package main

// 双指针，双循环， 模拟
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 || n == 0 {
		nums1 = append(nums1[:m], nums2...)
	} else {
		nums1 = append(nums1[:m], nums2...)
		for i := 0; i < n; i++ {
			for j := 0; j < m+i; j++ {
				if nums1[j] > nums2[i] {
					nums1[j], nums1[m+i] = nums1[m+i], nums1[j]
				}
			}
		}
	}

}
