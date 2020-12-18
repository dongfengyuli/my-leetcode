package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main(){
	//nums := []int{1,2,2,3,4,4,3}
	//tree := structures.Ints2TreeNode(nums)

	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	res := structures.BuildTree(preorder,inorder)
	fmt.Println(res)
	/*for res != nil{
		fmt.Println(res)
		fmt.Println("res.Val","res.Left","res.Right",res.Val,res.Left,res.Right)
	}*/
}

