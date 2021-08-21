package main

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func preorder(root *Node) (res []int) {
	if root == nil { // 别忘记判空
		return
	}
	res = append(res, root.Val)
	for _, v := range root.Children {
		res = append(res, preorder(v)...)
	}
	return
}
