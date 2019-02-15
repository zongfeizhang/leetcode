package tree

func FindFrequentTreeSum(root *TreeNode) []int {
	res := make([]int,0)
	record := make(map[int]int)
	maxTimes := 0
	_,maxTimes = dfsTreeSum(root,record,maxTimes)
	for key,val := range record{
		if val==maxTimes{
			res = append(res,key)
		}
	}
	return res
}

func dfsTreeSum(root *TreeNode,record map[int]int,maxTimes int)(int,int){
	if root==nil{
		return 0,maxTimes
	}else{
		leftSum,maxTimesLeft := dfsTreeSum(root.Left,record,maxTimes)
		rightSum,maxTimesRight := dfsTreeSum(root.Right,record,maxTimes)
		sum := leftSum + rightSum + root.Val
		record[sum] += 1
		maxTimes := max(record[sum],max(maxTimes,max(maxTimesLeft,maxTimesRight)))
		return sum,maxTimes
	}
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
