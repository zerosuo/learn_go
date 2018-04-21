package main

import "fmt"


// 阶乘 函数
func sum(n int) uint64 {

	var s uint64 = 1
	var sum uint64 = 0

	for i := 1; i <= n; i++ {
		// n的阶乘
		s = s * uint64(i)
		fmt.Printf("%d!=%v\n", i, s)
		// n的阶乘之和
		sum += s
	}
	return sum
}

func main() {
	var n int
	fmt.Println("请输入一个数字")
	fmt.Scanf("%d", &n)

	s := sum(n)
	fmt.Println(s)
}