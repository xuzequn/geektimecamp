package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//  递归， 分治
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

// 	if root == nil {
// 		return root
// 	}

// 	// 如果root 非空， root，查找左子树， 查找
// 	//  如果root 是 返回root
// 	if root.Val == p.Val || root.Val == q.Val{
// 		return root
// 	}
// 	// 没有就往左右左右子树找， 返回值
// 	left := lowestCommonAncestor(root.Left, p, q)
// 	right := lowestCommonAncestor(root.Right, p, q)

// 	// 如果左 右子树都找到
// 	if left != nil && right != nil{
// 		return root
// 	}

// 	// 如果只找到左子树
// 	if left!= nil {
// 		return left
// 	}

// 	return right

// }

// 记录父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	parent := map[int]*TreeNode{} // 记录元素的父亲节点
	visited := map[int]bool{}     // 父节点记录是否访问过

	//  记录父亲节点
	parent[root.Val] = nil
	var markparent func(root *TreeNode)
	markparent = func(root *TreeNode) {
		if root != nil {
			if root.Left != nil {
				parent[root.Left.Val] = root
				markparent(root.Left)
			}
			if root.Right != nil {
				parent[root.Right.Val] = root
				markparent(root.Right)
			}
		}
	}

	//  记录p节点的所有父亲节点。
	// 有问题的函数
	// var visit func(p *TreeNode) *TreeNode
	// visit = func(p *TreeNode) *TreeNode {
	// 	if visited[p.Val] {
	// 		return p
	// 	}
	// 	// 有父节点 ，标记,再访问父节点的父节点
	// 	visited[p.Val] = true
	// 	if v, ok := parent[p.Val]; ok {
	// 		if v != nil {
	// 			visit(v)
	// 		}
	// 	}
	// 	return nil
	// }

	markparent(root)

	for key := range parent {
		visited[key] = false
	}
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	// visit(p)
	// res := visit(q)
	return nil

}
