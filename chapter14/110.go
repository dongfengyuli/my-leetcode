package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{3,9,20,NULL,NULL,15,7,24,30}
	tree := structures.Ints2TreeNode(nums)
	res := structures.IsBalanced(tree)
	fmt.Println(res)
	//fmt.Println()
}
