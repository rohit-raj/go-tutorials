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
	a := 5
	fmt.Println(fact(a))
}