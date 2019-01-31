package dp

func IsScramble(s1 string, s2 string) bool {
	records := make(map[string]int)
	return dfs(s1,s2,records)
}

//recursive
func dfs(start,target string,records map[string]int) bool {
	if len(start)!=len(target){
		return false
	}
	mergeStr := start + target
	val,ok := records[mergeStr]
	if ok{
		if val==0{
			return false
		}
		return true
	}

	if start==target{
		records[mergeStr] = 1
		return true
	}


	///分割位置
	for i:=1;i<len(start);i++{
		//如果位置不交换
		leftRes := dfs(start[:i],target[:i],records)
		rightRes := dfs(start[i:],target[i:],records)
		if leftRes && rightRes{
			records[mergeStr] = 1
			return true
		}
		//如果交换位置
		leftRes = dfs(start[i:],target[:len(start)-i],records)
		rightRes = dfs(start[:i],target[len(start)-i:],records)
		if leftRes && rightRes{
			records[mergeStr] = 1
			return true
		}
	}
	records[mergeStr] = 0
	return false
}

