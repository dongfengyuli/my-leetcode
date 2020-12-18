package main

import (
	"fmt"
	"my-leetcode/structures"
)
var NULL = -1 << 63

func main(){
	num1 := []int{1,3,2,5}
	t1 := structures.Ints2TreeNode(num1)

	num2 := []int{2,1,3,NULL,4,7}
	t2 := structures.Ints2TreeNode(num2)


	res := structures.MergeTrees(t1,t2)
	fmt.Println(res)
	/*for res != nil{
		fmt.Println(res)
		fmt.Println("res.Val","res.Left","res.Right",res.Val,res.Left,res.Right)
	}*/
}