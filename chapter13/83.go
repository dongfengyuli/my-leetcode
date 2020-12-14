package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main(){
	nums := []int{1,2,2,2,12,16,19,19,36}
	head := structures.Ints2List(nums)

	res := structures.DeleteDuplicates(head)
	fmt.Println(res)
	//fmt.Println()

	for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}
}
