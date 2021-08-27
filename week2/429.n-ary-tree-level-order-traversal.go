package main

import (
	"container/list"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {

	//  用队列进行BFS ， 一个节点队列，按层序进入，一个valuelist 存储值。
	// 使用list 来模拟队列

	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	result := [][]int{}

	for queue.Len() > 0 {
		// queue pop , root.children pushback
		flowlist := []int{}
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Front()
			queue.Remove(node)
			for _, child := range node.Value.(*Node).Children {
				queue.PushBack(child)
			}
			flowlist = append(flowlist, node.Value.(*Node).Val)
		}
		result = append(result, flowlist)
	}
	return result

}
