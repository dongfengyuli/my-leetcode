package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{4,2,7,1,3,6,9}
	tree := structures.Ints2TreeNode(nums)
	res := structures.InvertTree(tree)
	fmt.Println(res)

}

