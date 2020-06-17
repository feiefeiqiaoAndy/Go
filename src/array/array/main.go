package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}

	//给一个数组找出的最大值

	var intArr = [5]int{1, -1, 9, 90, 11}

	maxValue := intArr[0]
	maxIndex := 0

	for i := 0; i < len(intArr); i++ {
		if maxValue < intArr[i] {
			maxValue = intArr[i]
			maxIndex = i
		}
	}
	fmt.Println(maxIndex)
	//一个数组求和  平均值
	var arr1 = [5]int{2,2,2,2,3}
	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Printf("sum=%v 平均值=%v\n", sum, sum/len(arr1))
	fmt.Println(sum / len(arr1))

	for _, val := range arr1 {
		sum += val
	}
	fmt.Printf("sum=%v 平均值=%v\n", sum,  float64(sum) / float64(len(arr1)))

	//数组反转

	var  op  [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <len(op); i++ {
		op[i]=rand.Intn(100)
	}
	fmt.Println(op)

	temp  := 0 //做一个临时变量
	for i := 0; i <len(op) /2 ; i++ {
		temp = op[len(op)-1 -i]      //最后一个元素
		op[len(op)-1 -i] = op[i]      //第一个
		op[i] = temp

		//可以不用temp  其实一样
		//op[i] = op[len(op)-1 -i]      //最后一个元素
		//op[len(op)-1 -i] = op[i]      要注销temp

	}


	fmt.Println(op)

}
