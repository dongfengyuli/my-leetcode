package main

import "fmt"

func main(){
	nums := []int{5,7,7,8,8,10}
	target := 8
	res := searchRange(nums,target)
	fmt.Println(res)
}


func searchRange(nums []int,target int) []int{
	if len(nums) == 0{
		return []int{-1,-1}
	}

	l,r := 0,len(nums)-1
	for l <= r{
		mid := (l+r)/2
		if nums[mid] == target{
			l1,r1 := mid,mid
			for l1>=0 && nums[l1] == target{
				l1--
			}
			for r1 <= len(nums) -1 && nums[r1]== target{
				r1++
			}
			return []int{l1+1,r1+1}
		}

		if target > nums[mid]{
			l = mid+1
		}else{
			r = mid-1
		}
	}
	return []int{-1,-1}
}