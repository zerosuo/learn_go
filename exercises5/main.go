package main

import (
	"time"
	"fmt"
)

func test() {
	time.Sleep(time.Millisecond*100)
}

func main() {

	// 这里注意的亮点是 格式化必须是2006/01/02 15：04：05
	now := time.Now()
	fmt.Println(now.Format("2006/01/02/ 15:04:05"))

	// 获取当前开始时间  获取结束时间
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Printf("const:%d us\n", (end - start)/1000)
}
