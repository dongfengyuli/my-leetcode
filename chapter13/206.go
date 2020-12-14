package main

import (
	"fmt"
	"my-leetcode/structures"
)

func main(){
	nums := []int{1,2,3,4,5,6}
	head := structures.Ints2List(nums)
	fmt.Println(head)

	res := structures.ReverseList(head)
	fmt.Println(res)
}

/*type ListNode struct {
	     Val int
	     Next *ListNode
}
*/
