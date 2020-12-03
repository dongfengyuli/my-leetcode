package main

import "fmt"

func main(){
	word1 := "sea"
	word2 := "eat"
	res := minDistance(word1,word2)
	fmt.Println(res)
}


func minDistance(word1 string, word2 string) int {
	return  len(word1) + len(word2) - longestCommonSubsequence(word1,word2)*2
}

//1143. 最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int,len(text1) + 1)
	for i := 0;i <= len(text1);i++ {
		dp[i] = make([]int,len(text2) + 1)
	}

	for i := 1;i <= len(text1);i++ {
		for j := 1;j <= len(text2);j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}