package dp

import (
	"sort"
)

func PermuteUnique(nums []int) [][]int {
	//排序
	//sort
	sort.Ints(nums)
	res := make([][]int,0)
	endFlag := false
	for endFlag==false{
		arr := make([]int,0)
		arr = append(arr,nums...)
		res = append(res,arr)
		endFlag = NextPermutation(nums)
	}
	return res
}

func NextPermutation(nums []int)bool{
	cur := len(nums)-1
	pre := cur - 1
	for pre>=0{
		if nums[pre]<nums[cur]{
			pos := find(nums,cur,len(nums)-1,nums[pre])
			swap(nums,pre,pos)
			reverse(nums,cur,len(nums)-1)
			return false
		}
		cur = pre
		pre = cur - 1
	}
	reverse(nums,0,len(nums)-1)
	return true
}

func find(nums []int,left,right,target int)int{
	mid := (left + right) / 2
	if nums[mid]>target && (mid+1==len(nums) || nums[mid+1]<=target){
		return mid
	}
	if nums[mid]<=target{
		return find(nums,left,mid-1,target)
	}else{
		return find(nums,mid+1,right,target)
	}
}

func reverse(nums []int,left,right int){
	for left<=right{
		swap(nums,left,right)
		left++
		right--
	}
}

func swap(nums []int,left,right int){
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}
