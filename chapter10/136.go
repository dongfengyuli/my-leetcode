package main

import "fmt"

func main(){
	nums := []int{4,1,2,1,2}
	res := singleNumber(nums)
	fmt.Println(res)
}


func singleNumber(nums []int) int{
	ans := 0
	for _,value := range nums{
		ans ^= value
	}
	return ans
}