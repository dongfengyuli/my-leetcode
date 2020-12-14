package main

import "fmt"

func main(){
	nums := []int{100,4,200,1,3,2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}


func longestConsecutive(nums []int) int{
	numSet := map[int]bool{}
	//numSet := make(map[int]string,0)
	for _,num := range nums{
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet{
		if !numSet[num-1]{
			currentNum := num
			currentSteak := 1
			for numSet[currentNum+1]{
				currentNum++
				currentSteak++
			}
			if longestStreak < currentSteak{
				longestStreak = currentSteak
			}
		}
	}
	return longestStreak
}