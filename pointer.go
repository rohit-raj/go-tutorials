package main

import "fmt"

func main(){
	x := 10
	pointer := &x //memory address of x

	fmt.Println(pointer)
	fmt.Println(*pointer) // printing value from the address

	*pointer = 20 //made changes in value stored at the address of x
	fmt.Println(x) // changed the value of x

	*pointer = *pointer**pointer // multiplication of values on the address of x
	fmt.Println(x)
	fmt.Println(*pointer)
}