package main

import "fmt"

func main(){
	nums := []int{2,5,6,0,0,1,2}
	target := 3

	res := search(nums,target)
	fmt.Println(res)
}


func search(nums []int,target int) bool{
	l,r := 0,len(nums) - 1
	for l <= r{
		mid := (l+r)/2
		if nums[mid] == target{
			return true
		}else if nums[l] == nums[mid]{
			l++
			continue
		}else if(nums[l] < nums[mid]){
			if nums[l] <= target && nums[mid] > target{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}else{
			if nums[r] >= target && nums[mid] < target{
				l = mid + 1
			}else{
				r = mid - 1
			}
		}
	}
	return false
}