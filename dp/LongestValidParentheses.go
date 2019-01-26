package dp

const(
	LEFT = "("
	RIGHT = ")"
)

type Ele struct{
	number int
	data string
}

func LongestValidParentheses(s string) int {
	if len(s)==0{
		return 0
	}
	partners := getPartner(s)
	dp := make([]int,len(s))
	for i:=len(s)-2;i>=0;i--{
		partnerIndex := partners[i]
		if string(s[i])==RIGHT || partnerIndex==0{
			dp[i] = 0
		}else if partnerIndex==len(s)-1{
			dp[i] = partnerIndex - i + 1
		}else{
			useCurPre := partnerIndex - i + 1
			dp[i] = max(useCurPre+dp[partnerIndex+1],dp[i+1])
		}
	}

	maxCount := dp[0]
	for _,val := range dp{
		maxCount = max(maxCount,val)
	}
	return maxCount
}


func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}


func getPartner(s string) []int{
	parteners := make([]int,len(s))
	stack := make([]Ele,0)
	for index,value := range s{
		ele := Ele{number:index,data:string(value)}
		if ele.data == LEFT{
			stack = append(stack,ele)
		}else{
			if len(stack)==0{
				parteners[index] = 0 //it has no partner
			}else{
				top := stack[len(stack)-1]
				parteners[top.number] = index
				stack = stack[:len(stack)-1]
			}
		}
	}
	return parteners
}



