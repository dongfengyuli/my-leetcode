package main

import "fmt"

func main()  {
	res := mySqrt(15)
	fmt.Println(res)
}



func mySqrt(x int) int{
	l,r := 0,x
	ans := -1
	for l <= r{
		mid := (l+r)/2
		if mid*mid <= x{
			ans = mid
			l = mid + 1
		}else{
			r = mid - 1
		}
	}
	return ans
}