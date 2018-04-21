package main

import (
	"fmt"
	"math"
)

// 输入一个区间范围输出其中的质数。

// 判断是否质数函数  bool返回值

func isPrime(n int) bool {

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// 定义2个 int 类型 a, b
	var a, b int


	fmt.Println("请输入区间范围的整数：")

	// 接收传入的2个数
	fmt.Scanf("%d%d", &a, &b)

	// 遍历 传入的2个值范围 每个数都传入 isPrime函数中 判断是否质数
	for i := a; i < b; i++ {

		// 如果是true 就打印这个数 并且中断 if这一层继续判断下一个
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}