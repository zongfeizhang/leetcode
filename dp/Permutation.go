package dp

func Permute(nums []int) [][]int {
	dp := initDpPermute(nums)
	for i:=1;i<=len(nums);i++{
		if len(dp[i-1])==0{
			dp[i] = [][]int{{nums[i-1]}}
		}else{
			for _,arr := range dp[i-1]{
				for index := range arr{
					if index==0{
						perm := make([]int,0)
						perm = []int{nums[i-1]}
						perm = append(perm,arr[0:]...)
						dp[i] = append(dp[i],perm)
					}
					if index>0 && arr[index-1]!=nums[i-1]{
						perm := make([]int,0)
						perm = append(perm,arr[0:index]...)
						perm = append(perm,nums[i-1])
						perm = append(perm,arr[index:]...)
						dp[i] = append(dp[i],perm)
					}
					//最后一个位置
					if index==len(arr)-1 && arr[index]!=nums[i-1]{
						perm := make([]int,0)
						perm = append(perm,arr[0:]...)
						perm = append(perm,nums[i-1])
						dp[i] = append(dp[i],perm)
					}

				}
			}
		}
	}
	return dp[len(nums)]
}

func initDpPermute(nums []int)[][][]int{
	dp := make([][][]int,len(nums)+1)
	return dp
}
