package main

import "container/heap"

type Iheap [][2]int

func (h Iheap) Less(i, j int) bool  { return h[i][1] < h[j][1] } // 小根堆
func (h Iheap) Len() int            { return len(h) }
func (h Iheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Iheap) Push(x interface{}) { (*h) = append((*h), x.([2]int)) }
func (h *Iheap) Pop() interface{}   { res := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return res }

func topKFrequent(nums []int, k int) []int {

	// 标记
	mark := map[int]int{}
	for i := 0; i < len(nums); i++ {
		mark[nums[i]]++
	}

	// 堆化
	h := Iheap{}
	heap.Init(&h)
	for i, v := range mark {
		heap.Push(&h, [2]int{i, v})

		if h.Len() > k {
			heap.Pop(&h)
		}

	}

	// 生成答案
	res := make([]int, k)
	for i := 0; i < k; i++ {

		res[k-i-1] = heap.Pop(&h).([2]int)[0]

	}

	return res
}
