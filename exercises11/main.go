package main

import "fmt"

func add(a, b int) int {
	return  a + b
}

func test() {
	return
}

func main() {
	c := add
	sum := c(200, 300)
	fmt.Println(sum)

}

