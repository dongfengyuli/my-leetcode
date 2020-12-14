package main

import "fmt"

func main(){
	nums := []int{2, 7, 11, 15}
	target := 13
	res := twoSum(nums,target)
	fmt.Println(res)
}


func twoSum(nums []int,target int) []int{
	m := make(map[int]int,len(nums))
	for i,b := range nums{
		if j,ok := m[target-b];ok{
			return []int{j,i}
		}
		m[nums[i]] = i
	}
	return nil
}