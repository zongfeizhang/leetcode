package dp

type Node struct{
	left,right *Node
	value string
}

func IsScramble1(s1 string, s2 string) bool {
	root := buildTree(s1)
	if len(s1)!=len(s2){
		return false
	}
	if root==nil{
		if s2!=""{
			return false
		}else{
			return true
		}
	}
	//带有记录的dfs
	record := make(map[string]int)
	return dfs1(root,s2,record)
}

func dfs1(root *Node,target string,record map[string]int) bool{
	mergeStr := root.value + target
	val,ok := record[mergeStr]
	if ok{
		if 1==val{
			return true
		}
		return false
	}
	//这个节点可能交换也可能不交换
	if isLeaf(root){
		return root.value==target
	}
	mid := len(target) / 2
	leftTarget := target[:mid]
	rightTarget := target[mid:]
	if dfs1(root.left,leftTarget,record) && dfs1(root.right,rightTarget,record){
		record[mergeStr] = 1
		return true
	}

	leftTarget = target[:len(root.right.value)]
	rightTarget = target[len(root.right.value):]
	if dfs1(root.right,leftTarget,record) && dfs1(root.left,rightTarget,record){
		record[mergeStr] = 1
		return true
	}
	record[mergeStr] = 0
	return false
}

func isLeaf(node *Node) bool{
	return node.left==nil && node.right==nil
}

func buildTree(str string)*Node{
	if len(str)==0{
		return nil
	}
	if len(str)==1{
		node := &Node{value:str}
		return node
	}
	//分成两份
	mid := len(str) / 2
	leftStr := str[:mid]
	rightStr := str[mid:]
	root := &Node{value:str}
	root.left = buildTree(leftStr)
	root.right = buildTree(rightStr)
	return root
}