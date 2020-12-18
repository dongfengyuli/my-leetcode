package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{4,1,6,0,2,5,7,NULL,NULL,NULL,3,NULL,NULL,NULL,8}
	tree := structures.Ints2TreeNode(nums)
	res := structures.ConvertBST(tree)
	fmt.Println(res)

}

