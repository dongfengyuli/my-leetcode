package main

import "fmt"

func main(){
	nums := []int{3,0,1}
	res := missingNumber(nums)
	fmt.Println(res)
}


func missingNumber(nums []int) int {
	sum := len(nums)
	for i:=0; i < len(nums);i++{
		sum += i - nums[i]
	}
	return sum
}


/*func missingNumber(nums []int) int {
	res	:=	len(nums)
	for i,num:=range nums{
		res^=i^num
	}
	return res
}*/

