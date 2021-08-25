package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		order(node.Left)            //左
		res = append(res, node.Val) // 中
		order(node.Right)           // 右

	}
	order(root)

	return

}
