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
	a := 10
	fmt.Println(fibo(a))
}