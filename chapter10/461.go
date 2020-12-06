package main

import "fmt"

func main(){
	x := 1
	y := 4
	res := hammingDistance(x,y)
	fmt.Println(res)
}


func hammingDistance(x int, y int) int{
	//两个整数之间的汉明距离是对应位置上数字不同的位数。
	//计算 x 和 y 之间的汉明距离，可以先计算 x XOR y，然后统计结果中等于 1 的位数
	n := x ^ y
	ans := 0
	for n != 0{
		n &= n - 1
		ans++
	}
	return ans
}