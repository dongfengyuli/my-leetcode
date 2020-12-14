package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main()  {
	nums1 := []int{4,1,8,4,5}
	list1 := structures.Ints2List(nums1)
	nums2 := []int{5,0,1,8,4,5}
	list2 := structures.Ints2List(nums2)

	res := structures.GetIntersectionNode(list1,list2)
	fmt.Println(res)

	for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}
}
