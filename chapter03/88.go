package main

import "fmt"

func main(){
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3

	merge(nums1,m,nums2,n)
	fmt.Println(nums1)

}

func merge(nums1 []int,m int,nums2 []int,n int){
	left,right,p := m-1,n-1,m+n-1
	for left >= 0 && right >= 0{
		if nums1[left] > nums2[right]{
			nums1[p] = nums1[left]
			left--
		}else{
			nums1[p] = nums2[right]
			right--
		}
		p--
	}
	//当num2未完全填充的时候
	for right >= 0{
		nums1[p] = nums2[right]
		p--
		right--
	}
}