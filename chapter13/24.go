package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main()  {
	nums := []int{1,2,4,6,12,16,19,25,36}
	head := structures.Ints2List(nums)

	res := structures.SwapPairs(head)
	fmt.Println(res)

	for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}
}
