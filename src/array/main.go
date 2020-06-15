package main

import "fmt"

func main() {

	var list = [4]int{1,2,3,4}
	fmt.Println(list)

	for i := 0; i <100 ; i++ {
		fmt.Println(i)
	}

	var names = [5]string{"a","b","c","d","f"}

	//range
	for k,v :=range names {
		fmt.Println(k,v)
	}

	for _ ,v :=range names {
		fmt.Println(v)
	}
}
