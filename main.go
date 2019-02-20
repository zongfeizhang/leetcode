package main

import (
	"leetcode/array"
	"fmt"
)

func main() {
	nums := []int{0,0,0,0}
	target := 0
	res := array.FourSum(nums,target)
	fmt.Println(res)
}