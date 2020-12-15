package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{1,NULL,6,3,5,7,13}
	tree := structures.Ints2TreeNode(nums)
	res := structures.PostOrderTraversal(tree)
	fmt.Println(res)
}



