package main

import "fmt"

func main(){
	res := countPrimes(20)

	fmt.Println(res)
}


func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}


//我们考虑这样一个事实：如果 xx 是质数，那么大于 xx 的 xx 的倍数 2x,3x,\ldots2x,3x,… 一定不是质数，因此我们可以从这里入手。

