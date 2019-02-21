package array

func LengthOfLongestSubstring(s string) int {
	set := make(map[string]int)
	maxLen := 0
	left := 0
	right := left
	for left<=right && right<len(s){
		ch := string(s[right])
		times := set[ch]
		if times==0{
			set[ch]++
			if right==len(s)-1{
				maxLen = max(maxLen,right-left+1)
			}
			right++
		}else{
			curLen := right - left
			maxLen = max(maxLen,curLen)
			for set[ch]==1{
				set[string(s[left])]--
				left++
			}
		}
	}
	return maxLen
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}