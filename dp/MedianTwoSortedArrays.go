package dp

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	allLen := len(nums1) + len(nums2)
	medians := make([]int,0)
	mid := allLen/2
	medians = append(medians,mid)
	if allLen%2==0{
		medians = append(medians,mid-1)
	}
	res := 0.0
	for _,median := range medians{
		result := binarySearch(nums1,0,len(nums1)-1,nums2,0,len(nums2)-1,median)
		if result==-1{
			result = binarySearch(nums2,0,len(nums2)-1,nums1,0,len(nums1)-1,median)
		}
		res += result
	}
	return res/float64(len(medians))
}


func binarySearch(a []int,leftA,rightA int,b []int,leftB,rightB, target int)float64{
	if leftA>rightA{
		return -1
	}

	mid := (leftA + rightA) / 2
	pos := findLastLessThan(b,leftB,rightB,a[mid]) //mid会不会超范围
	leftCount := mid + pos + 1
	if target==leftCount{
		return float64(a[mid])
	}else if target==leftCount-1 && pos!=-1 && b[pos]==a[mid]{
		return float64(b[pos])
	}else{
		if target<leftCount{
			if pos==-1{
				return float64(a[target])
			}
			return binarySearch(a,leftA,mid-1,b,leftB,pos-1,target)
		}else{
			return binarySearch(a,mid+1,rightA,b,pos+1,rightB,target)
		}
	}
}

func findLastLessThan(arr []int,left,right,value int)int{
	if left>right{//首尾边界
		return right
	}
	mid := (left+right) / 2
	if arr[mid]>value{
		return findLastLessThan(arr,left,mid-1,value)
	}else{
		if mid+1>=len(arr)||arr[mid+1]>value{
			return mid
		}else{
			return findLastLessThan(arr,mid+1,right,value)
		}
	}
}







