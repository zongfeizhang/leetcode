package array

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var minDiff float64
	var res []int
	for i:=0;i<len(nums)-2;i++{
		left := i+1
		right := len(nums)-1
		for left<right{
			curSum := nums[i] + nums[left] + nums[right]
			curDiff := math.Abs(float64(curSum-target))
			if res==nil{
				res = []int{nums[i],nums[left],nums[right]}
				minDiff = curDiff
			}

			res,minDiff = update(nums,i,left,right,res,curDiff,minDiff)

			if curSum<target{
				left++
			}else{
				right--
			}
		}
	}
	return res[0]+res[1]+res[2]
}

func update(nums []int,i,left,right int,res []int,curDiff,minDiff float64)([]int,float64){
	if curDiff<minDiff{
		return []int{nums[i],nums[left],nums[right]},curDiff
	}
	return res,minDiff
}