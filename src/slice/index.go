package main

import "fmt"

func main() {

	a:=[10]int{1,2,3,4,5,6,7,8,9,10}
	/*
	slice:=arr[]
	[start,end]
	arr[:end]
	arr[start:]
	 */
	fmt.Println("-----------对应的是索引-----------")
	s1 :=a[:5]  //
	//s2 :=a[3:5]  //
	//s3 :=a[5:]  //
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(s3)

	fmt.Println(a)
	fmt.Println("s1",s1)
	//扩容就是新开辟一块空间
	s1 = append(s1,99,88,77,22,33,100,200,300,400,500,600,700,800,900)
	fmt.Println(a)  //[1 2 3 4 5 6 7 8 9 10]

	fmt.Println(s1)  //[1 2 3 4 5 99 88 77 22 33 100 200 300 400 500 600 700 800 900]







}
