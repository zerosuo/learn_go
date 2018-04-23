package main

import (
	"math/rand"
	"fmt"

)

func main() {
	// 先随机从100产生一个数
	n := rand.Intn(100)

	// for死循环 让保证一直循环直到符合条件
	for {
		// 输入的数字
		var input int
		fmt.Scanf("%d\n", &input)
		// 默认设置数字不对
		flag := false
		switch {
		// 如果相等 就返回true
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("less" )
		}
		// 检测到如果flag 是true了
		if flag {
			// 执行终止
			break
		}

	}
}
