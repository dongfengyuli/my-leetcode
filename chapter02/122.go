package main

import "fmt"

func main(){
	prices := []int{7,1,5,3,6,4}
	res := maxProfit(prices)
	fmt.Println(res)
}



func maxProfit(prices []int) int{
	if len(prices) < 1{
		return 0
	}

	ans := 0
	for i:=1;i<len(prices);i++{
		if prices[i]>prices[i-1]{
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}