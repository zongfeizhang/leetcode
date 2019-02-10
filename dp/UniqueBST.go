package dp

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
 }


func GenerateTrees(n int) []*TreeNode {
	dp := make([][][]*TreeNode,n+1) //节点个数 根+节点个数
	for i:=0; i<n+1;i++{
		dp[i] = make([][]*TreeNode,n+1)
	}

	for i:=1;i<len(dp);i++{//根
		for j:=i;j<len(dp[i]);j++{//节点个数
			dp[i][j] = make([]*TreeNode,0)
			//left
			leftArr := make([]*TreeNode,0)
			for left:=1;left<i;left++{
				leftArr = append(leftArr,dp[left][i-1]...)
			}
			//right
			rightArr := make([]*TreeNode,0)
			for right:=i+1;right<=j;right++{
				rightArr = append(rightArr,dp[right][j-i]...)
			}

			//汇总处理
			if len(leftArr)==0 && len(rightArr)==0{
				node := &TreeNode{Val:i}
				dp[i][j] = append(dp[i][j],node)
			}
			//左边为空
			if len(leftArr)==0{
				for _,child := range rightArr{
					node := &TreeNode{Val:i}
					node.Right = child
					dp[i][j] = append(dp[i][j],node)
				}
			}
			//右边为空
			if len(rightArr)==0{
				for _,child := range leftArr{
					node := &TreeNode{Val:i}
					node.Left = child
					dp[i][j] = append(dp[i][j],node)
				}
			}

			//两边都不为空
			for _,leftChild := range leftArr{
				for _,rightChild := range rightArr{
					node := &TreeNode{Val:i}
					node.Left = leftChild
					node.Right = rightChild
					dp[i][j] = append(dp[i][j],node)
				}
			}
		}
	}

	res := make([]*TreeNode,0)
	for i:=1;i<len(dp);i++{
		res = append(res,dp[i][n]...)
	}
	return res
}




