package main

import (
	"fmt"
)

func main() {
	var myName string      //declare types
	myName = "Hello"       //can change value
	yourName := "Timepass" //declare new variable with value using :=
	fmt.Println(myName + " " + yourName)
	//print type using %T
	fmt.Printf("%T\n", myName) //works only with Printf
	var x = 23                 //gives int
	fmt.Printf("%T\n", x)
	var y int64 = 344577777777777777 //explicit type declaration
	fmt.Printf("%T\n", y)
	var u = 45.0
	fmt.Printf("%T\n", u) //gives float64
	//var j uint16=-2 usnsigned cant be negative as it gives overflow error
	// int8,16,32,64
	// unint 8,16,32,64
	//float,complex-32,64
	var k byte = 15
	fmt.Printf("%T\n", k) //gives uint8
}
