package bs

func FindMin(nums []int) int {
	if nums[0]<=nums[len(nums)-1]{
		return nums[0]
	}
	return bst(nums,0,len(nums)-1)
}

func bst(nums []int,left,right int)int{
	mid := (left + right) / 2
	//判断mid所处位置
	if nums[mid]>=nums[left] && nums[mid]>=nums[right]{//qian
		if nums[mid]>nums[mid+1]{
			return nums[mid+1]
		}
		return bst(nums,mid+1,right)
	}else{
		if nums[mid]<nums[mid-1]{
			return nums[mid]
		}
		return bst(nums,left,mid-1)
	}
}
