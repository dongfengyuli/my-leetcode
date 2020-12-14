package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main(){
	nums := []int{1,2,3,4,5}
	head := structures.Ints2List(nums)

	res := structures.RemoveNthFromEnd(head,2)
	fmt.Println(res)

	for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}
}