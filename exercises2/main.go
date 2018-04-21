package main

import "fmt"

// 判断是否水仙花
func isNumber(n int) bool {
	// 取模 获得 三位数的每个位上的数
	i := n % 10
	j := (n / 10) % 10
	k := (n / 100) % 10

	sum := i*i*i + j*j*j + k*k*k
	if sum == n {
		return  true
	}
	return  false
}

func main() {
	var n, m int

	fmt.Println("请输入两个数字中间空格")
	// 获取n 和m
	fmt.Scanf("%d%d", &n, &m)

	// 判断一下 n, m是否是数字
	if n == 0 && m == 0 {
		fmt.Println("您输入有误，请输入两个数字。")
	}

	for i := n; i < m; i++ {
		// 传入核心判断是否水仙花函数
		if isNumber(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}