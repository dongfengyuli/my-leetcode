package main

import "fmt"

func main()  {
	ratings := []int{1,2,2,7,4,3,9}
	res := candy(ratings)

	fmt.Println(res)
}

func candy(ratings []int) int{
	can := make([]int,len(ratings))
	for i:=0;i<len(ratings);i++{
		can[i] = 1
		if i == 0{
			continue
		}
		if ratings[i] > ratings[i-1] && can[i] <= can[i-1]{
			can[i] = can[i-1] + 1
		}
	}

	for i := len(ratings) - 2;i >=0;i--{
		if ratings[i] > ratings[i+1] && can[i] <= can[i+1]{
			can[i] = can[i+1] + 1
		}
	}

	sum := 0
	for i:=0;i<len(can);i++ {
		sum += can[i]
	}

	return sum
}