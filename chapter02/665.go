package main

import "fmt"

func main(){
	nums := []int{8,5,7,3,1}

	res := checkPossibility(nums)
	fmt.Println(res)
}


func checkPossibility(nums []int) bool{
	if nil == nums || len(nums) < 0{
		return true
	}

	if len(nums) == 1{
		return true
	}

	count := 0
	for i:=1;i<len(nums) && count <=1;i++{
		if nums[i-1] > nums[i]{
			count++
			if i-2 < 0 || nums[i-2] <= nums[i]{
				nums[i-1] = nums[i]
			}else{
				nums[i] = nums[i-1]
			}
		}
	}
	return count <= 1
}