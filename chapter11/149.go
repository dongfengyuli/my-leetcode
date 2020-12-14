package main

import "fmt"

func main(){
	points := [][]int{
		{1,1},
		{3,2},
		{5,3},
		{4,1},
		{2,3},
		{1,4},
	}

	res := maxPoints(points)
	fmt.Println(res)
}


//斜率，用最简约分子分母表示，为防止溢出，类型为int64
type K struct {
	m int64 //分子
	n int64 //分母
}

//对每个点，检查其他点与该点的斜率
func maxPoints(points [][]int) int {
	n := len(points)
	if n<3 {return n}

	res := 0
	same := K{0,0}

	for i:=0;i<n-1;i++ {
		hmap := map[K]int{}
		temp := 0
		for j:=i+1;j<n;j++ {
			k := getK(points[i],points[j])
			hmap[k]++
			if k!=same {
				temp = max(temp,hmap[k])
			}
		}
		//此时，从i点出发，处于同一直线上的点最多有temp+1个
		//再加上与i点相同的点的个数hmap[same]即可
		//更新全局答案res
		res = max(res,temp+1+hmap[same])
	}
	return res
}


//根据两点求斜率
func getK(p1, p2 []int) K {
	dx := int64(p1[0]-p2[0])
	dy := int64(p1[1]-p2[1])
	if dx==0&&dy==0 {return K{int64(0),int64(0)}} //p1,p2是同一点
	if dx==0 {return K{int64(1),int64(0)}} //p1,p2斜率为无穷
	if dy==0 {return K{int64(0),int64(1)}} //p1,p2斜率为0
	if dx<0 { //统一设置dx为正
		dx=-dx
		dy=-dy
	}
	d := gcd(dx,dy)
	return K{dx/d,dy/d} //返回最简约之比
}

//最大公约数
func gcd(a, b int64) int64 {
	if a==0 {return b}
	return gcd(b%a,a)
}

func max(a, b int) int {
	if a>b {return a}
	return b
}

