package main

import "fmt"

func print(n int) {
	// 循环行数
	for i := 0 ; i < n+1 ; i++ {
		// 打印A
		for m := 0 ; m < i; m++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func main() {
	print(9)
}

// 一直模块化思维