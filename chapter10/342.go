package main

import "fmt"

func main(){
	res := isPowerOfFour(134)
	fmt.Println(res)
}

func isPowerOfFour(num int) bool{
	if num < 0{
		return false
	}
	for num > 10{
		flag := num % 4
		num /= 4
		if flag != 0{
			return false
		}
	}

	if num == 1 || num == 4{
		return true
	}
	return false
}