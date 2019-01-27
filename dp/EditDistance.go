package dp

import "math"

func MinDistance(word1 string, word2 string) int {
	if word1==""{
		return len(word2)
	}
	if word2==""{
		return len(word1)
	}


	dp := make([][]int,len(word1)+1)
	initDP(dp,len(word2)+1)

	for i:=0;i<len(word2);i++{
		dp[0][i] = i
	}
	for i:=0;i<len(word1);i++{
		dp[i][0] = i
	}
	for i:=1;i<=len(word1);i++{
		start := string(word1[i-1])
		for j:=1;j<=len(word2);j++{
			end := string(word2[j-1])
			//第一种情况
			if start==end{
				dp[i][j] = min(dp[i-1][j-1],dp[i][j])
			}else{
				dp[i][j] = min(dp[i-1][j-1]+1,dp[i][j])
			}
			//第二种情况
			dp[i][j] = min(dp[i][j],dp[i][j-1]+1)
			//第三种情况
			dp[i][j] = min(dp[i][j],dp[i-1][j]+1)
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}


func initDP(dp [][]int,size int){
	for index := range dp{
		dp[index] = make([]int,size)
		for col:=0;col<len(dp[index]);col++{
			dp[index][col] = math.MaxInt16
		}
	}
}
