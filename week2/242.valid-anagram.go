package main

// // 哈希map
// func isAnagram(s string, t string) bool {

// 	var ms map[rune]int
// 	ms = make(map[rune]int)

// 	var mt map[rune]int
// 	mt = make(map[rune]int)

// 	if len(s) != len(t) {
// 		return false
// 	}

// 	for _, v := range s {
// 		ms[v]++
// 	}

// 	for _, v := range t {
// 		mt[v]++
// 	}

// 	return reflect.DeepEqual(ms, mt) // golang 两个map 都比较，用reflact.DeepEqual, 不确定对象类型都可以用这个，效率有点慢。

// }

// 排序
// func isAnagram(s string, t string) bool {

// 	sb, tb := []byte(s), []byte(t)

// 	// sort.Slice 的用法
// 	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
// 	sort.Slice(tb, func(i, j int) bool { return tb[i] < tb[j] })

// 	return string(sb) == string(tb)
// }

// 哈希 2
func isAnagram(s string, t string) bool {

	var sc, tc [26]int

	for i := 0; i < len(s); i++ {
		sc[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		tc[t[i]-'a']++
	}

	return sc == tc

}
