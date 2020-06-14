package main

import "fmt"

func main() {
	var  arr1   [3]int
	arr1[0] = 5
	arr1[1] = 10
	arr1[2] = 15
	fmt.Println(arr1)

	var  list    = [2]string{"hello","world"}
	fmt.Println(list)

	arr := [4]int{100,200,300,400}
	fmt.Println(arr)



	var names = [...]string{"a","b","c"}
	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	//range 循环
	var str = [...]string{"aa","bb","cc"}
	for k, v := range str {
		println(k,v)

	}

	// 注意不要用  fmt.Printf() 还是输出数组
}
