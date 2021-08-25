package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	mp := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		// 排序
		list := []byte(strs[i])
		sort.Slice(list, func(j, k int) bool { return list[j] < list[k] })
		strlist := string(list)
		mp[strlist] = append(mp[strlist], strs[i])

	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
