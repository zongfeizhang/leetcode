package tree

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	res := make([][]int,0)
	if root==nil{
		return res
	}

	for len(queue)!=0{
		arr := fill(queue)
		res = append(res,arr)
		childs := bornChilds(queue)
		queue = childs
	}
	return res
}

func bornChilds(parent []*TreeNode)[]*TreeNode{
	childs := make([]*TreeNode,0)
	for _,node := range parent{
		if node.Left != nil{
			childs = append(childs,node.Left)
		}
		if node.Right != nil{
			childs = append(childs,node.Right)
		}
	}
	return childs
}

func fill(t []*TreeNode)[]int{
	arr := make([]int,len(t))
	for index,node := range t{
		arr[index] = node.Val
	}
	return arr
}
