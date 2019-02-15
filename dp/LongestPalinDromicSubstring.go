package dp

func LongestPalindrome(s string) string {
	dp := initDp(s)
	longest := ""
	for i:=1; i<len(dp);i++{
		for j:=1; j<=i;j++{
			if s[i-1]==s[j-1] && (j+1>i-1||dp[j+1][i-1]==1){
				dp[j][i] = 1
				curStr := s[j-1:i]
				if len(longest)<len(curStr){
					longest = curStr
				}
			}
		}
	}
	return longest
}

func initDp(s string)[][]int{
	dp := make([][]int,len(s)+1)
	for i := range dp{
		dp[i] = make([]int,len(s)+1)
		dp[i][i] = 1
	}
	return dp
}