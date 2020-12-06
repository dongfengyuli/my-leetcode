package main

import "fmt"

func main(){
	res := trailingZeroes(25)

	fmt.Println(res)
}

func trailingZeroes(n int) int{
	count := 0
	for n > 0{
		count += n/5
		n = n/5
	}
	return count
}