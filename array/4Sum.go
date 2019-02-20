package array

import "sort"

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int,0)
	for i:=0;i<len(nums)-3;i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		remain := target - nums[i]
		arrs := threeSum(nums,i+1,remain)
		for _,arr := range arrs{
			newArr := []int{nums[i]}
			newArr = append(newArr,arr...)
			res = append(res,newArr)
		}
	}
	return res
}

func threeSum(nums[]int,start, target int)[][]int{
	res := make([][]int,0)
	for i:=start;i<len(nums)-2;i++{
		if i>start && nums[i]==nums[i-1]{
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left<right{
			curSum := nums[i] + nums[left] + nums[right]
			if curSum==target{
				res = append(res,[]int{nums[i],nums[left],nums[right]})
				left++
				for left<right && nums[left]==nums[left-1]{
					left++
				}
				right--
				for left<right && nums[right]==nums[right+1]{
					right--
				}
			}
			if curSum<target{
				left++
			}
			if curSum>target{
				right--
			}
		}
	}
	return res
}
