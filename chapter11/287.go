package main

import "fmt"

func main(){
	nums := []int{3,1,3,4,2}

	res := findDuplicate(nums)
	fmt.Println(res)
}


func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
		//fmt.Println("slow",slow)
		//fmt.Println("fast",fast)
	}

	//fmt.Println("-------")
	//fmt.Println("slow",slow)
	//fmt.Println("fast",fast)

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}