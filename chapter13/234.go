package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main()  {
	nums := []int{1,2,4,2,1}
	head := structures.Ints2List(nums)

	res := structures.Ispalindrome(head)
	fmt.Println(res)


	/*for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}*/
}
