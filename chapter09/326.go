package main

import "fmt"

func main(){
	res := isPowerOfThree(26)
	fmt.Println(res)
}



func isPowerOfThree(n int) bool {
	return n > 0 && (0 == 1162261467%n)
}


// 1162261467 is 3^19,  3^20 is bigger than int