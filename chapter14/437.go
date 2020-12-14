package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main()  {
	nums := []int{10,5,-3,3,2,NULL,11,3,-2,NULL,1}
	tree := structures.Ints2TreeNode(nums)
	res := structures.PathSum(tree,8)
	fmt.Println(res)

}
