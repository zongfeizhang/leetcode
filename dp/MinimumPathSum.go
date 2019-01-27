package dp

func minPathSum(grid [][]int) int {
	if grid==nil||len(grid)==0||grid[0]==nil{
		return 0
	}
	dp := make([][]int,len(grid))
	for row:=0;row<len(dp);row++{
		dp[row] = make([]int,len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	//0列
	for row:=1;row<len(grid);row++{
		dp[row][0] = dp[row-1][0] + grid[row][0]
	}
	//0行
	for col:=1;col<len(grid[0]);col++{
		dp[0][col] = dp[0][col-1] + grid[0][col]
	}

	for row:=1;row<len(grid);row++{
		for col:=1;col<len(grid[0]);col++{
			dp[row][col] = min(dp[row-1][col],dp[row][col-1]) + grid[row][col]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}