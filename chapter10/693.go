package main


func main(){

}


/*func hasAlternatingBits(n int) bool {
	for n != 0 {
		if (n & 1) == ((n >> 1) & 1) {
			return false // 相邻相等
		}
		n >>= 1
	}
	return true
}
*/


func hasAlternatingBits(n int) bool {
	res := decimalToBinary(n)
	for i, temp := 1,res[0];i < len(res);i++ {
		if temp == res[i] {
			return false
		}else{
			temp = res[i]
		}
	}
	return true
}


func decimalToBinary(n int) []int {
	if 0 == n {
		return []int{0}
	}

	ans := make([]int,0)
	for 0 != n {
		ans = append(ans,n%2)
		n /= 2
	}
	return ans
}