package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{3,2,3,4,5}
	tree := structures.Ints2TreeNode(nums)
	res := structures.DiameterOfBinaryTree(tree)
	fmt.Println(res)
}
