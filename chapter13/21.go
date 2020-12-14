package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main()  {
	nums1 := []int{1,2,4}
	list1 := structures.Ints2List(nums1)
	nums2 := []int{1,3,5}
	list2 := structures.Ints2List(nums2)

	res := structures.MergeTwoLists(list1,list2)
	fmt.Println(res)
	for nil != res{
		fmt.Println("res.Val","res.Next",res.Val,res.Next)
		res = res.Next
	}
}