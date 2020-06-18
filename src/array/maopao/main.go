package main

import "fmt"

func maopao(arr *[5]int) {
	temp := 0
	for j := 0; j < len(*arr)-1; j++ {
		for i := 0; i < len(*arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
	}
	fmt.Println(*arr)
}

func main() {
	var arr = [5]int{5, 88, 4, 66, 33}

	maopao(&arr)

}
