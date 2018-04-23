package main

import "fmt"


func modify(p *int) {
	fmt.Println(p)
	*p = 1000900
	return
}

func main() {

	// 打印a的地址
	a := 10
	fmt.Println(&a)

	// 打印 a的变量的地址
	var p *int
	p = &a

	fmt.Println("the address of p:", &p)
	fmt.Println("the value of p", p)

	// 打印p指向a地址变量的值
	fmt.Println("the value of p point to variable:", *p)

	fmt.Println(*p) // 10

	// 修改了*p的值
	*p = 100
	fmt.Println(a) // 100

	b := 999

	p = &b
	*p = 5

	fmt.Println(a) // 100
	fmt.Println(b) // 5

	modify(&a) // 地址
	fmt.Println(a) // 1000900
}
