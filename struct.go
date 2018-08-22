package main

import "fmt"

type human struct {
	bone uint
	blood float64
	nose uint
	eyes uint
}

func main(){
	boy1 := human{bone : 206,
				blood : 5.5,
				nose : 1,
				eyes : 2}

	boy2 := human{206, 5.5, 1, 2}

	fmt.Println("boy 1 :", boy1.bone)
	fmt.Println("boy 2 :", boy2.blood)
}