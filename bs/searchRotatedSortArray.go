package bs

func Search(nums []int, target int) bool {
	return bstRotatedSortArray(nums,0,len(nums)-1,target)
}


func bstRotatedSortArray(nums []int,left,right,target int)bool{
	if left>right{
		return false
	}
	mid := (left + right) / 2
	if nums[mid]==target{
		return true
	}else if nums[mid]==nums[left] && nums[mid]==nums[right]{
		return bstRotatedSortArray(nums,left,mid-1,target) || bstRotatedSortArray(nums,mid+1,right,target)
	}else{
		//left,right组成的区间位于有序部分
		if nums[mid]>=nums[left] && nums[mid]<=nums[right]{
			if nums[mid]<target{
				return bstRotatedSortArray(nums,mid+1,right,target)
			}else{
				return bstRotatedSortArray(nums,left,mid-1,target)
			}
		}
		//[mid]位于前半部分
		if nums[mid]>=nums[left] && nums[mid]>=nums[right]{
			if nums[mid]>target && nums[right]<target{//
				return bstRotatedSortArray(nums,left,mid-1,target)
			}else{
				return bstRotatedSortArray(nums,mid+1,right,target)
			}
		}
		//[mid]位于后半部分
		if nums[mid]<=nums[left] && nums[mid]<=nums[right]{
			if nums[mid]<target && target<=nums[right]{
				return bstRotatedSortArray(nums,mid+1,right,target)
			}else{
				return bstRotatedSortArray(nums,left,mid-1,target)
			}
		}
		return false
	}
}