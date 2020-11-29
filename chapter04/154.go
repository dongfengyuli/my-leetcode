package main

import "fmt"

func main(){
	nums := []int{4,5,6,7,0,1,2}
	res := findMin(nums)
	fmt.Println(res)
}

func findMin(nums []int) int{
	l,r := 0,len(nums) - 1
	for l < r{
		mid := (l + r)/2
		if nums[mid] < nums[r]{
			r = mid
		}else if nums[mid] > nums[r]{
			l = mid + 1
		}else{
			r--
		}
	}
	return nums[l]
}