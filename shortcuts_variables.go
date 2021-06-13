package main

import (
	"fmt"
)

func main() {
	//mass decalaration :)
	var (
		name     = "Hrishikesh"
		age      = 21
		language = "Python"
	)
	fmt.Println(name, age, language)
	//declare many together
	character, number := "H", 30
	fmt.Println(character, number)
	//string formatting
	fmt.Printf("%v is your score out of %v\n", 30, "100")
}
