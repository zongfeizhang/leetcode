package tree

func smallestFromLeaf(root *TreeNode) string {
	if root==nil{
		return ""
	}else{
		leftStr := smallestFromLeaf(root.Left)
		rightStr := smallestFromLeaf(root.Right)
		ch := getCh(root.Val)
		if root.Left==nil{
			return rightStr+ch
		}
		if root.Right==nil{
			return leftStr+ch
		}
		if leftStr<rightStr{
			return leftStr+ch
		}else{
			return rightStr+ch
		}
	}
}

func getCh(val int)string{
	return string('a'+val)
}
