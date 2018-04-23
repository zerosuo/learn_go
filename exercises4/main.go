package main

import (
	"fmt"
	"strings"
	"strconv"
)

func urlProcess(url string) string {

	result := strings.HasPrefix(url, "http://")
	if !result{
		url = fmt.Sprintf("http://%s", url)
	}

	return url

}

func pathProcess(path string) string {

	result := strings.HasPrefix(path, "/")
	if !result{
		path = fmt.Sprintf("%s", path)
	}

	return path

}

func main() {
	var (
		url string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)

	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)

	// 替换字符串
	str := "   hello world abc \n"
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println("replace:", result)

	// 字符出现次数
	count := strings.Count(str, "l")
	fmt.Println("count:", count)

	// 重复三次字符串
	result = strings.Repeat(str, 3)
	fmt.Println("repeat:", result)

	// 全转小写
	result = strings.ToLower(str)
	fmt.Println("lower:", result)

	// 全转大写
	result = strings.ToUpper(str)
	fmt.Println("upper", result)

	// 去掉首尾字符
	result = strings.TrimSpace(str)
	fmt.Println("trimSpace:", result)

	// 去掉字符串首尾cut字符
	result = strings.Trim(str, "\n\r")
	fmt.Println("trim:", result)

	// 去掉尾cut字符
	result = strings.TrimLeft(str, "\n\r")
	fmt.Println("trimLeft:", result)

	// 返回字符以空格的切片
	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	//返回split切割的字符串
	splitResult = strings.Split(str, "l")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	//用l把所有字符串链接起来
	str2 := strings.Join(splitResult, "l")
	fmt.Println("join:", str2)

	//把整数i转为字符串
	str2 = strconv.Itoa(1000)
	fmt.Println("itoa:", str2)

	// 把一个字符串转成整数
	number, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("can not convert to int,", err)
		return
	}
	fmt.Println("number:", number)

}
