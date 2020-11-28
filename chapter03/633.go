package main

import (
	"fmt"
	"math"
)

func main(){
	res := judgeSquareSum(10)
	fmt.Println(res)
}

func judgeSquareSum(c int) bool{
	if c < 0{
		return false
	}

	left,right := 0,int(math.Sqrt(float64(c)))
	for left <= right{
		current := left*left + right*right
		if current < c{
			left++
		}else if current > c{
			right--
		}else{
			return true
		}
	}
	return false
}