package tree

func pruneTree(root *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	left := pruneTree(root.Left)
	right := pruneTree(root.Right)
	if root.Val==0&&left==nil&&right==nil{
		return nil
	}
	root.Left = nil
	root.Right = nil
	if left!=nil{
		root.Left = left
	}
	if right!=nil{
		root.Right = right
	}
	return root
}
