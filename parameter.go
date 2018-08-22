package main

import "fmt"

/**
 * sum is a function which takes :
 * x of type float64
 * y of type float64
 * and returns value of type float64
 * this can be rewritten as
 * func sum(x, y float64) float64 {}
 */
func sum1(x float64, y float64) float64 {
	return x+y
}

func sum2(x, y float64) float64 {
	return x+y
}

func multipleReturn(a,b string)(string,string){
	return a,b
}

/**
 * sum is a function which takes :
 * a of type float64
 * b of type float64
 * this can be rewritten as
 * var a, b float64 = 2.3, 9.9
 */
func main(){
	var a float64 = 2.3
	var b float64 = 9.9
	fmt.Println("sum1 is ", sum1(a,b))

	var c, d float64 = 2.3, 9.9
	fmt.Println("sum2 is ", sum2(c,d))

	num1, num2 := 2.3, 9.9
	fmt.Println("sum3 is ", sum2(num1, num2))

	w1, w2 := "Hi", "Rohit"
	fmt.Println(multipleReturn(w1,w2))
}