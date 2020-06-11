package main

import "fmt"

func main() {

	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	var (
		a string
		b int
		c bool
		d float64
	)
	fmt.Println(a, b, c, d)

	var app string = "类型"
	var time int = 15
	fmt.Println(app, time)

	//类型推导
	var name1 = "andy"
	var n = 8
	fmt.Println(name1, n)

	//短变量声明
	number := 10
	fmt.Println(number)

}
