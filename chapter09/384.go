package main

import (
	"fmt"
	"math/rand"
)

func main(){
	nums := []int{1,5,7,11,14,6}
	s := Constructor(nums)
	res := s.Shuffle()
	fmt.Println("s.nums",s.nums)
	fmt.Println("res",res)
}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
