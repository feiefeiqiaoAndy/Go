package main

func main() {

	var  i  int  =10

	var  ptr  *int =  &i    //地址

	 *ptr   =  100           // 1 ptr = &i 就是去i的地址  2 *ptr 就拿到i的地址进行赋值100

	println(*ptr , &ptr)  //  *ptr的值100   ptr 本身也有一个地址

	/*
	结果  100 0xc000035f70
	 */
}
