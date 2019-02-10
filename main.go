package main

import (
	"fmt"
	"leetcode/dp"
)

func main() {
	//s := "()(()"
	//res := dp.LongestValidParentheses(s)
	//fmt.Println(res)
	//
	//arr := [][]int{
	//	{1, 3, 1},
	//	{1, 5, 1},
	//	{4, 2, 1},
	//}
	//res = dp.MinPathSum(arr)
	//fmt.Println(res)
	///////
	//word1 := "intention"
	//word2 := "execution"
	//res = dp.MinDistance(word1, word2)
	//fmt.Println(res)
	//
	//fmt.Println("=========================")
	//s1 := "abcde"
	//s2 := "caebd"
	//flag := dp.IsScramble(s1, s2)
	//fmt.Println(flag)
	//
	//result := dp.GenerateTrees(3)
	//count := len(result)
	//fmt.Println(count)
	//
	//fmt.Println("=============结果============")
	n := 4
	fmt.Println(dp.NumTrees(n))
}