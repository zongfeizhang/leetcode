package bs

func SearchRange(nums []int, target int) []int {
	left := bstLeft(nums,0,len(nums)-1,target)
	right := bstRight(nums,0,len(nums)-1,target)
	res := []int{left,right}
	return res
}

//最左边
func bstLeft(nums []int,left,right,target int)int{
	if left>right{
		return -1
	}
	mid := (left + right) / 2
	if nums[mid]==target && (mid==0||nums[mid-1]!=target){
		return mid
	}
	if nums[mid]<target{
		return bstLeft(nums,mid+1,right,target)
	}else{
		return bstLeft(nums,left,mid-1,target)
	}
}

//最右边
func bstRight(nums []int,left,right,target int)int{
	if left>right{
		return -1
	}
	mid := (left + right) / 2
	if nums[mid]==target && (mid==len(nums)-1 || nums[mid+1]!=target){
		return mid
	}
	if nums[mid]>target{
		return bstRight(nums,left,mid-1,target)
	}else{
		return bstRight(nums,mid+1,right,target)
	}
}
