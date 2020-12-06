package main

import "fmt"

func main(){
 	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
 	res := dailyTemperatures(T)
 	fmt.Println(res)
}


func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int,length)
	stack := []int{}

	for i:=0;i<length;i++{
		temperature := T[i]
		for len(stack) > 0 && temperature > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack,i)
	}
	return ans
}

