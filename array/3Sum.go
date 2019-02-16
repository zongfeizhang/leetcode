package array

import "sort"

func ThreeSum(nums []int) [][]int {
	res := make([][]int,0)
	m := make(map[int]int)
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if i<2{
			m[nums[i]]++
			continue
		}
		for j:=i-1;j>=0;j--{
			//去重
			if (j+1<=i-1&&nums[j]==nums[j+1])||m[nums[i]]>2 || (m[nums[i]]==2&&nums[i]!=0){
				continue
			}
			//向前扫描 已经以[i-1]为根扫过了
			if m[nums[i]]>0 && nums[j]!=nums[i]{
				continue
			}

			cur := 0 - nums[i] - nums[j]
			if ((cur!=nums[j]&&m[cur]>0)||m[cur]>1)&&cur<=nums[j]{
				arr := []int{nums[i],nums[j],cur}
				res = append(res,arr)
			}
		}
		m[nums[i]]++
	}
	return res
}
