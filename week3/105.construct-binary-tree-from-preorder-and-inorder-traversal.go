package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	// 创建跟节点
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	// 中序遍历root的index
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}

	}

	// 创建 左子树 和右子树

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root

}
