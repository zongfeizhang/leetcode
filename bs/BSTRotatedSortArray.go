package bs

func search(nums []int, target int) int {
	return binarySearch(nums,0,len(nums)-1,target)
}

func binarySearch(nums []int,left,right,target int)int{
	if left>right{
		return -1
	}else{
		mid := (left + right) / 2
		if nums[mid]==target{
			return mid
		}else if nums[mid]<target{
			//位于前半段还是后半段
			if (nums[mid]>=nums[left] && nums[mid]<=nums[right]) || (nums[mid]>=nums[left] && nums[mid]>=nums[right]) || target<=nums[right]{//连续有序
				return binarySearch(nums,mid+1,right,target)
			}
			return binarySearch(nums,left,mid-1,target)
		}else{
			if (nums[mid]>=nums[left] && nums[mid]<=nums[right]) || (nums[mid]<=nums[left] && nums[mid]<=nums[right]) || target>nums[right]{//连续有序
				return binarySearch(nums,left,mid-1,target)
			}
			return binarySearch(nums,mid+1,right,target)
		}
	}
}