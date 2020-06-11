package main
import "fmt"

// const (
// 	pi = 3.14
// 	e = 2
// )
// const (
// 	n1 = 10 
// 	n2 = 10 
// 	n3 = 10 
// )
// const (
// 	n1 = 10 
// 	n2 
// 	n3 
// )
// 每增加一行+1
//  const (
// 	 n1 = iota
// 	 n2 = iota
// 	 n3 = iota
// 	 n4 = iota
//  )
//忽略
//  const (
// 	n1 = iota
// 	n2 
// 	_
// 	n4 
// )
//插队
const (
	n1 = iota  //0
	n2  = iota  //1
	n3 =100   //100
	n4  = iota   // 3 
)
const n5 = iota    //0
func main ()  {
	//常量
	const pi = 3.1415
	const e = 2.7
	fmt.Println(n1,n2,n3,n4,n5)
}
