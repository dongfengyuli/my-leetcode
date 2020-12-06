package main

import "fmt"

func main(){
	nums := []int{4,3,2,7,8,2,3,1}
	res := findDisappearedNumbers(nums)

	fmt.Println(res)
}



func findDisappearedNumbers(nums []int) []int{
	hash := make(map[int]int)
	res := []int{}

	for _,i := range nums{
		hash[i] = 1
	}

	for j:=1;j<=len(nums);j++{
		if hash[j] != 1{
			res = append(res,j)
		}
	}
	return res
}