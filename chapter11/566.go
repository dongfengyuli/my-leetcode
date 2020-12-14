package main

import "fmt"

func main(){
	nums := [][]int{
		{1,4},
		{2,5},
		{6,15},
		{8,25},
	}
	res := matrixReshape(nums,2,4)
	fmt.Println(res)
}


func matrixReshape(nums [][]int, r int, c int) [][]int {
	if nil == nums || 0 == len(nums) || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	temp := make([]int,0)
	for row := 0;row < len(nums);row++ {
		for col := 0;col < len(nums[0]);col++ {
			temp = append(temp,nums[row][col])
		}
	}

	fmt.Println("temp",temp)

	ans := make([][]int,r)
	for row := 0;row < r;row++ {
		tmp := make([]int,c)
		for col := 0;col < c;col++ {
			tmp[col] =temp[row*c + col]
		}
		ans[row] = tmp
	}
	return ans
}