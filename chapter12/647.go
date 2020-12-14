package main

import "fmt"

func main(){
	s := "abcdefe"
	res := countSubstrings(s)
	fmt.Println(res)
}



func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2 * n - 1; i++ {
		l, r := i / 2, i / 2 + i % 2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

/*方法一：中心拓展
思路与算法

计算有多少个回文子串的最朴素方法就是枚举出所有的回文子串，而枚举出所有的回文字串又有两种思路，分别是：

1.枚举出所有的子串，然后再判断这些子串是否是回文；
2.枚举每一个可能的回文中心，然后用两个指针分别向左右两边拓展，当两个指针指向的元素相同的时候就拓展，否则停止拓展。
*/
