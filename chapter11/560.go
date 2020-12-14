package main

import "fmt"

func main(){
	nums := []int{3,4,7,2,-3,1,4,2}
	res := subarraySum(nums,7)
	fmt.Println(res)
}


//前缀和 + 哈希表优化

func subarraySum(nums []int, k int) int {
	count,pre := 0,0
	m := map[int]int{}
	m[0] = 1

	for i:=0;i<len(nums);i++ {
		pre += nums[i]
		if _,ok := m[pre-k];ok {
			count += m[pre - k]
		}
		m[pre] += 1
	}

	//fmt.Println("m",m)
	return count
}