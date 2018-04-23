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

	// 转化字符串
	str := "   hello world abc \n"
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println("replace:", result)

	// 字符出现次数
	count := strings.Count(str, "l")
	fmt.Println("count:", count)

	// 查找出现三次的字符
	result = strings.Repeat(str, 3)
	fmt.Println("repeat:", result)

	//
	result = strings.ToLower(str)
	fmt.Println("lower:", result)

	//
	result = strings.ToUpper(str)
	fmt.Println("upper", result)

	//
	result = strings.TrimSpace(str)
	fmt.Println("trimSpace:", result)

	//
	result = strings.Trim(str, "\n\r")
	fmt.Println("trim:", result)

	//
	result = strings.TrimLeft(str, "\n\r")
	fmt.Println("trimLeft:", result)

	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(str, "l")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "l")
	fmt.Println("join:", str2)

	str2 = strconv.Itoa(1000)
	fmt.Println("itoa:", str2)

	number, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println("can not convert to int,", err)
		return
	}
	fmt.Println("number:", number)

}
