package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main(){
	nums := []int{2,1,3,5,6,4,7,19,36}
	head := structures.Ints2List(nums)

	res := structures.OddEvenList(head)
	fmt.Println(res)

	for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}
}
