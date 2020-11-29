package main

import "fmt"

func main(){
	nums := []int{1,1,2,3,3,4,4,8,8}
	res := singleNonDuplicate(nums)

	fmt.Println(res)
}


func singleNonDuplicate(nums []int) int{
	l,r := 0,len(nums)
	for l < r{
		mid := (l+r)/2
		// 如果中间索引不是偶数，我们可以对其 -1，只针对偶数位置进行判断
		if mid %2 == 1{
			mid--
		}
		// 中间索引是偶数，那么如果它前面不存在只出现一次的数，当前值应该是相同元素的第一个
		if nums[mid] == nums[mid + 1]{
			l = mid + 2
		}else{
			// 否则只出现一次的数在前面
			r = mid
		}
	}
	return nums[l]
}
