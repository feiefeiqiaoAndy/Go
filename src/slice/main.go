package main

import "fmt"

func main() {

	n :=[]int{1,2}
	fmt.Println(n)

   list :=	make([]int,3,8)

   fmt.Printf("容量:%d,长度:%d\n",cap(list),len(list))
	list = append(list,5 )
	list = append(list,5,55,66,77,88,99,100,200,300,400 )

	//把n的数组 append list数组
	list = append(list,n...)

   //fmt.Println(list)
	for i:=0;i < len(list);i++ {
		fmt.Println(list[i])
	}
	s1 :=[]int{1,2}
	fmt.Print(s1)
	s1 = append(s1, 3,4)
	fmt.Printf("%p\n",s1)
	s1 = append(s1, 30,40)
	fmt.Printf("%p\n",s1)

	//地址改变了
}
