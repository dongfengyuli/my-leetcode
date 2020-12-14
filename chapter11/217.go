package main

import "fmt"

func main(){
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	res := containsDuplicate(nums)
	fmt.Println(res)
}

func containsDuplicate(nums []int) bool{
	dict := make(map[int]int,len(nums))

	for _,value := range nums{
		dict[value]++
	}

	for _,v := range dict{
		if v > 1{
			return true
		}
	}
	return false
}