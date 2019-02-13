package main

import (
	"leetcode/bs"
	"fmt"
)

func main() {
	nums := []int{5,7,7,8,8,10}
	res := bs.SearchRange(nums,6)
	fmt.Println(res)
}