package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{1,2,2,3,4,4,3}
	tree := structures.Ints2TreeNode(nums)
	res := structures.IsSymmetric(tree)
	fmt.Println(res)
}

