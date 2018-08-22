package main

import "fmt"

func fact(x int) int {
	if(x <= 1){
		return 1
	} else {
		return x * fact(x-1)
	}
}

func main(){
	var a int
	fmt.Println("Enter a number to print the factorial")
	fmt.Scanf("%d", &a)
	fmt.Println(fact(a))
}