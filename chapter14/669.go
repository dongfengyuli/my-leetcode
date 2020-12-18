package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{3,0,4,NULL,2,NULL,NULL,1}
	tree := structures.Ints2TreeNode(nums)
	res := structures.TrimBST(tree,1,3)
	fmt.Println(res)
}
