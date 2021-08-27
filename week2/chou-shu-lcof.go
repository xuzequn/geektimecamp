package main

import "container/heap"

type theap []int

func (h theap) Less(i, j int) bool { return h[i] < h[j] }
func (h theap) Len() int           { return len(h) }
func (h theap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// *h 解了 *theap 类型的指针，也就是获取到了h地址存放的值，取值操作
func (h *theap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *theap) Pop() interface{}   { res := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return res }

func nthUglyNumber(n int) int {

	fector := []int{2, 3, 5}
	ugly := -1
	h := theap{}
	// 初始化堆
	heap.Init(&h)
	heap.Push(&h, 1)
	// 会重复，所有记录
	visit := map[int]bool{}
	for i := 0; i < n; i++ {
		if h[0] == ugly {
			heap.Pop(&h)
		}
		ugly = heap.Pop(&h).(int)
		for _, v := range fector {
			re := ugly * v
			if _, ok := visit[re]; !ok {
				visit[re] = true
				heap.Push(&h, re)
			}
		}
	}
	return ugly

}
