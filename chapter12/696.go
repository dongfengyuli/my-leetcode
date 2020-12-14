package main

import "fmt"

func main(){
	s := "00110011"
	res := countBinarySubstring(s)

	fmt.Println(res)
}


func countBinarySubstring(s string) int{
	preRun,curRun,ans := 0,1,0
	for i:=1;i<len(s);i++{
		if s[i] == s[i-1]{
			curRun++
		}else{
			preRun = curRun
			curRun = 1
		}
		if preRun >= curRun{
			ans++
		}
	}
	return ans
}
