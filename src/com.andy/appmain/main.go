package main

import (
	"com.andy/services"
	"fmt"
)

func main()  {

	result := services.GetUser()
	fmt.Println(result)
}
