package main

import (
	"strconv"
	"fmt"
)

func number(a string) int {
	number, err := strconv.Atoi(a)
	if err != nil {
		println("转化出错")
	}
	return number

}

func main() {
	var m string
	fmt.Scanf("%s", &m)

	a := number(m)
	fmt.Println(a)

}
// 第一次执行结果： 转化出错 0
// 仔细一看 传值 m 没有 传到 &m 的地址
