package main

import "fmt"

func main(){
	flowerbed := []int{1,0,0,0,1}
	n := 2
	res := canPlaceFlowers(flowerbed,n)
	fmt.Println(res)
}


func canPlaceFlowers(flowerbed []int,n int) bool{
	if n > len(flowerbed){
		return false
	}
	blank := 0
	for i:=0;i<len(flowerbed);i++{
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed) -1 || flowerbed[i+1] == 0){
			flowerbed[i] =1
			blank++
		}
	}
	return blank >= n
}