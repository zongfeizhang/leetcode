package dp

func IsScramble(s1 string, s2 string) bool {
	if len(s1)!=len(s2){
		return false
	}else{
		dp := buildDP(s1,s2)
		//起始值特殊处理
		for i:=0;i<len(s1);i++{
			for j:=0;j<len(s2);j++{
				if s1[i]==s2[j]{
					dp[i][j][1] = true
				}
			}
		}

		///根据状态转移方程慢慢推
		for k:=2;k<len(s1)+1;k++{
			for i:=0;i+k<=len(s1);i++{
				for j:=0;j+k<=len(s2);j++{
					//分割位置
					for s:=1;s<k;s++{
						if (dp[i][j][s]==true&&dp[i+s][j+s][k-s]==true)||
							(dp[i][j+k-s][s]==true&&dp[i+s][j][k-s]==true){
							dp[i][j][k] = true
						}
					}
				}
			}
		}
		return dp[0][0][len(s1)]
	}
}

func buildDP(s1,s2 string)[][][]bool{
	dp := make([][][]bool,len(s1))
	for i := range dp{
		dp[i] = make([][]bool,len(s2))
		for index := range dp[i]{
			dp[i][index] = make([]bool,len(s2)+1)
		}
	}
	return dp
}
