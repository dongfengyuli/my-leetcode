package main

import (
	"fmt"
	"my-leetcode/structures"
)

var NULL = -1 << 63

func main(){
	nums := []int{3, 9, 20,NULL,11,15,7}
	tree := structures.Ints2TreeNode(nums)
	res := structures.AverageOfLevels(tree)
	fmt.Println(res)
}


