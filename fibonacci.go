package main

import "fmt"

func fibo(x int) int {
	if(x <= 2){
		return 1
	} else {
		return fibo(x-1) + fibo(x-2)
	}
}

func main(){
	var a int
	fmt.Println("Enter a number to print the fibonacci value")
	fmt.Scanf("%d", &a)
	fmt.Println(fibo(a))
}