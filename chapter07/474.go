package main

import "fmt"

func main(){
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	res := findMaxForm(strs,m,n)
	fmt.Println(res)
}


func findMaxForm(strs []string, m int, n int) int {
	statistic := func(str string) (int,int) {
		var zero,one = 0,0
		for _,char := range str {
			if char == '0'{
				zero++
			}else {
				one++
			}
		}
		return zero,one
	}
	dp := make([][]int,m+1)
	//初始化
	for i:=0;i<=m;i++{
		dp[i] = make([]int,n+1)
	}

	//进行动态规划
	for _,str := range strs {
		zero,one := statistic(str)
		//当zero或者one比m或者n大时，这个字符串就不用判断了。
		if zero > m || one > n {
			continue
		}
		for i:=m;i>=zero;i-- {
			for j:=n;j>=one;j-- {
				dp[i][j] = max(dp[i][j],dp[i-zero][j-one]+1)
			}
		}
	}

	fmt.Println("dp",dp)
	return dp[m][n]
}


func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}