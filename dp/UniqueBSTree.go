package dp

func numTrees(n int) int {
	records := make([]int,n+1)
	kinds := dfsKinds(n,records)
	return kinds
}

func dfsKinds(n int,records []int) int{
	if records[n]!=0{
		return records[n]
	}
	if n==0{
		return 0
	}
	if n==1{
		return 1
	}
	kinds := 0
	for i:=0;i<=n-1;i++{
		left := dfsKinds(i,records)
		right := dfsKinds(n-i-1,records)
		kinds += left * right
		if left*right==0{
			kinds += left + right
		}
	}
	records[n] = kinds
	return kinds
}