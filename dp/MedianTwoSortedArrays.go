package dp

import "math"

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
		result := kth(nums1,0,len(nums1),nums2,0,len(nums2),median)
		res += result
	}
	return res/float64(len(medians))
}

func kth(a []int,leftA,rightA int,b []int,leftB,rightB, k int)float64{
	if leftA==rightA{
		return float64(b[k])
	}
	if leftB==rightB{
		return float64(a[k])
	}
	midA := (leftA+rightA) / 2
	midB := (leftB+rightB) / 2
	if midA+midB+1==k{
		return math.Max(float64(a[midA]),float64(b[midB]))
	}else if midA+midB+1<k{
		if a[midA]>b[midB]{
			return kth(a,leftA,rightA,b,midB,leftB,k-midB-1)
		}else{
			return kth(a,midA,rightA,b,leftB,rightB,k-midA-1)
		}
	}else{
		if a[midA]>b[midB]{
			return kth(a,leftA,midA,b,leftB,rightB,k)
		}else{
			return kth(a,leftA,rightA,b,leftB,midB,k)
		}
	}
}







